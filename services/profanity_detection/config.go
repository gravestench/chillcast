package profanity_detection

import (
	"fmt"

	"github.com/gravestench/chillcast/services/config_file_manager"
)

const (
	cfgKeyProfanities    = "profanities (comma-delimited)"
	cfgKeyFalsePositives = "false-positives (comma-delimited)"
	cfgKeyFalseNegatives = "false-negatives (comma-delimited)"
)

func (s *Service) ConfigFileName() string {
	return "profanity_detection.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	g := cfg.Group(s.Name())

	g.Set(cfgKeyProfanities, "")
	g.Set(cfgKeyFalsePositives, "")
	g.Set(cfgKeyFalseNegatives, "")

	return
}
