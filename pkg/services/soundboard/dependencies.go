package soundboard

import (
	"github.com/gravestench/runtime/pkg"

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

func (s *Service) ResolveDependencies(runtime pkg.IsRuntime) {
	for _, service := range runtime.Services() {
		if candidate, ok := service.(config_file_manager.Manager); ok {
			s.cfgManager = candidate
		}
	}
}
