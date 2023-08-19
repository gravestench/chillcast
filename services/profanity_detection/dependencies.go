package profanity_detection

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/chillcast/services/config_file_manager"
)

func (s *Service) DependenciesResolved() bool {
	return s.cfgManager != nil
}

func (s *Service) ResolveDependencies(rt runtime.R) {
	for _, service := range rt.Services() {
		if candidate, ok := service.(config_file_manager.Dependency); ok {
			s.cfgManager = candidate
		}
	}
}
