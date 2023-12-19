package text_to_speech

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	if cfg, _ := s.Config(); cfg == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if candidate, ok := service.(config_file_manager.Manager); ok {
			s.cfgManager = candidate
		}
	}
}
