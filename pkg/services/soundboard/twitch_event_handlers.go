package soundboard

import (
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

type SoundBoardObject struct {
	name      string
	enabled   bool
	directory string
	files     struct {
		literals []string
		pattern  regexp.Regexp
	}
	triggers struct {
		literals []string
		pattern  regexp.Regexp
	}
	antispam struct {
		count    int
		interval time.Duration
	}
}

func (s *Service) getGroupsThatTriggerOnTwitchEvent(twitchEvent, message string) (triggerGroups []string) {
	cfg, err := s.Config()
	if err != nil {
		s.Logger().Error().Msgf("getting config: %v", err)
		return nil
	}

	for _, groupName := range cfg.GroupKeys() {
		sb := cfg.Group(groupName)

		events := sb.GetString(keyTwitchEvent)

		if !sb.GetBool(keyEnable) {
			continue
		}

		if !strings.Contains(events, twitchEvent) {
			continue
		}

		for _, literal := range sb.GetStrings(keyTriggerLiterals) {
			if strings.Contains(message, literal) {
				triggerGroups = append(triggerGroups, groupName)
				continue
			}
		}

		if regexp.MustCompile(sb.GetString(keyTriggerRegex)).MatchString(message) {
			triggerGroups = append(triggerGroups, groupName)
		}

		triggerGroups = append(triggerGroups, groupName)
	}

	return triggerGroups
}

func (s *Service) objectFromConfigFile(name string, cfg config_file_manager.Object) (obj SoundBoardObject) {
	obj.name = name
	obj.enabled = cfg.GetBool(keyEnable)
	obj.directory = expandHomeDirectory(cfg.GetString(keyAudioDirectory))

	obj.files.literals = cfg.GetStrings(keyAudioFilenamesLiterals)
	obj.files.pattern = *regexp.MustCompile(cfg.GetString(keyAudioFilenamesRegex))

	obj.triggers.literals = cfg.GetStrings(keyTriggerLiterals)
	obj.triggers.pattern = *regexp.MustCompile(cfg.GetString(keyTriggerRegex))

	obj.antispam.count = cfg.GetInt(keySpamCount)
	obj.antispam.interval = cfg.GetDuration(keySpamInterval)

	return
}

func (s *Service) pickSoundUsingObject(obj SoundBoardObject) (path string) {
	existingSoundFilePaths := make([]string, 0)

	fromLiterals, err := s.getExistingAbsoluteFilePaths(obj.directory, obj.files.literals)
	if err != nil {
		s.Logger().Error().Msgf("getting existing files from literals: %v", err)
	}

	fromRegex, err := s.getExistingAbsoluteFilePathsByRegex(obj.directory, obj.files.pattern)
	if err != nil {
		s.Logger().Error().Msgf("getting existing files from literals: %v", err)
	}

	existingSoundFilePaths = append(existingSoundFilePaths, fromLiterals...)
	existingSoundFilePaths = append(existingSoundFilePaths, fromRegex...)
	existingSoundFilePaths = s.removeDuplicates(existingSoundFilePaths)

	return pickRandom(existingSoundFilePaths)
}

func (s *Service) shouldTrigger(obj SoundBoardObject, triggerables ...string) bool {
	if !obj.enabled {
		return false
	}

	var shouldTrigger bool

	for _, triggerable := range triggerables {
		for _, literal := range obj.triggers.literals {
			if strings.Contains(triggerable, literal) {
				shouldTrigger = true
				break
			}
		}

		if obj.triggers.pattern.MatchString(triggerable) {
			shouldTrigger = true
			break
		}
	}

	if !shouldTrigger {
		return false
	}

	if _, found := s.antispam[obj.name]; !found {
		s.antispam[obj.name] = 0
	}

	// increment the count, wait the specified interval in a goroutine and
	// then decrement it
	s.antispam[obj.name] += 1
	go func(antispam map[string]int, name string, interval time.Duration) {
		s.Logger().Debug().Msgf("Antispam for %s: %d", name, antispam[name])
		time.Sleep(interval)
		antispam[name] -= 1
	}(s.antispam, obj.name, obj.antispam.interval)

	notSpamming := s.antispam[obj.name] <= obj.antispam.count

	return shouldTrigger && notSpamming
}

func (s *Service) OnTwitchPrivateMessage(message twitch.PrivateMessage) {
	triggerGroups := s.getGroupsThatTriggerOnTwitchEvent("OnPrivateMessage", message.Message)
	if len(triggerGroups) < 1 {
		return
	}

	randGroup := pickRandom(triggerGroups)

	cfg, err := s.Config()
	if err != nil {
		s.Logger().Error().Msgf("getting config: %v", err)
		return
	}

	g := cfg.Group(randGroup)
	obj := s.objectFromConfigFile(randGroup, g)
	s.pickSoundUsingObject(obj)

	if s.shouldTrigger(obj, message.Message) {
		go s.playAudio(s.pickSoundUsingObject(obj))
	}
}

func pickRandom(choices []string) string {
	if len(choices) < 1 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return choices[r.Intn(len(choices))]
}

// Returns a list of full absolute filepaths for the given filenames that exist in the parentDirectory.
func (s *Service) getExistingAbsoluteFilePaths(parentDirectory string, filenames []string) ([]string, error) {
	var validPaths []string

	for _, filename := range filenames {
		fullPath := filepath.Join(parentDirectory, filename)
		if _, err := os.Stat(fullPath); err == nil {
			validPaths = append(validPaths, fullPath)
		} else if !os.IsNotExist(err) {
			continue
			// return nil, err // some other error occurred
		}
	}

	return validPaths, nil
}

// Returns a list of full absolute filepaths for files in parentDirectory matching the regex pattern.
func (s *Service) getExistingAbsoluteFilePathsByRegex(parentDirectory string, pattern regexp.Regexp) ([]string, error) {
	var validPaths []string

	err := filepath.WalkDir(parentDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && pattern.MatchString(d.Name()) {
			validPaths = append(validPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return validPaths, nil
}

// Removes duplicates from a slice of strings.
func (s *Service) removeDuplicates(strings []string) []string {
	set := make(map[string]struct{})
	var result []string

	for _, str := range strings {
		if _, exists := set[str]; !exists {
			set[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func expandHomeDirectory(path string) string {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			path = strings.Replace(path, "~", homeDir, 1)
		}
	}
	return path
}
