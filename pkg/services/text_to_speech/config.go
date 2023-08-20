package text_to_speech

import (
	"fmt"
	"path/filepath"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

func (s *Service) ConfigFileName() string {
	return "text_to_speech.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	g := cfg.Group("Text to speech")

	cfgDir := s.cfgManager.ConfigDirectory()
	g.SetDefault("directory", filepath.Join(cfgDir, "audio_files"))
	g.SetDefault("mplayer-handler", false)

	return
}
