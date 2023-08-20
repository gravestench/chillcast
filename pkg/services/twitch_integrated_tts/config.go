package twitch_integrated_tts

import (
	"fmt"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

func (s *Service) ConfigFileName() string {
	return "twitch_integrated_text_to_speech.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	g := cfg.Group(s.Name())

	g.SetDefault("remove profanity", false)

	return
}
