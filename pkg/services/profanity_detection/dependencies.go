package profanity_detection

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

func (s *Service) DependenciesResolved() bool {
	return s.cfgManager != nil
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if candidate, ok := service.(config_file_manager.Dependency); ok {
			s.cfgManager = candidate
		}
	}
}
