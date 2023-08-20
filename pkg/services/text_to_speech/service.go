package text_to_speech

import (
	"os"
	"strings"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg"

	config_file_manager2 "github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

type Service struct {
	logger     *zerolog.Logger
	cfgManager config_file_manager2.Manager
	speech     htgotts.Speech
}

func (s *Service) Init(rt pkg.IsRuntime) {
	var cfg *config_file_manager2.Config

	for { // wait until the config or default config is saved + loaded
		time.Sleep(time.Second)

		if cfg, _ = s.Config(); cfg != nil {
			break
		}
	}

	g := cfg.Group("Text to speech")

	cfgDir := g.GetString("directory")

	var handler handlers.PlayerInterface = &handlers.Native{}
	if g.GetBool("mplayer-handler") {
		handler = &handlers.MPlayer{}
	}

	s.speech = htgotts.Speech{
		Folder:   expandHomeDirectory(cfgDir),
		Language: "en",
		Handler:  handler,
	}
}

func (s *Service) Name() string {
	return "Text-to-speech"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
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
