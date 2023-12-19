package twitch_integrated_tts

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
	"github.com/gravestench/chillcast/pkg/services/profanity_detection"
	"github.com/gravestench/chillcast/pkg/services/text_to_speech"
)

func (s *Service) DependenciesResolved() bool {
	if s.tts == nil {
		return false
	}

	if s.pd == nil {
		return false
	}

	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if candidate, ok := service.(text_to_speech.Dependency); ok {
			s.tts = candidate
		}

		if candidate, ok := service.(profanity_detection.Dependency); ok {
			s.pd = candidate
		}

		if candidate, ok := service.(config_file_manager.Dependency); ok {
			s.cfgManager = candidate
		}
	}
}
