package soundboard

import (
	"fmt"
	"strings"
	"time"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

const (
	keyEnable          = "enabled"
	keyTwitchEvent     = "twitch event"
	keyTriggerLiterals = "trigger literals"
	keyTriggerRegex    = "trigger regex"

	keyAudioDirectory         = "audio files directory"
	keyAudioFilenamesLiterals = "audio filename literals"
	keyAudioFilenamesRegex    = "audio filename regex"

	keySpamCount    = "spam count"
	keySpamInterval = "spam interval"
)

func (s *Service) ConfigFileName() string {
	return "soundboard.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	defaultAntispamInterval := time.Second * 30

	{
		g := cfg.Group("Example #1")

		g.Set(keyEnable, false)
		g.Set(keyTwitchEvent, "OnPrivateMessage")
		g.Set(keyTriggerLiterals, "fart")
		g.Set(keyTriggerRegex, "(butt|fart)")

		g.Set(keyAudioDirectory, `C:\path\to\audio\files`)
		g.Set(keyAudioFilenamesLiterals, `fart1.wav,fart2.mp3,fart3.mp3`)
		g.Set(keyAudioFilenamesRegex, `fart\d.(wav|mp3)`)

		g.Set(keySpamCount, 3)
		g.Set(keySpamInterval, defaultAntispamInterval.String())
	}

	{
		g := cfg.Group("_help")

		g.Set("twitch event list (these are the possible events you can use to trigger)", strings.Join([]string{
			"OnPrivateMessage",
		}, ","))

	}

	return
}
