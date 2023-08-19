package twitch_integrated_tts

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/chillcast/services/config_file_manager"
	"github.com/gravestench/chillcast/services/profanity_detection"
	"github.com/gravestench/chillcast/services/text_to_speech"
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

func (s *Service) ResolveDependencies(runtime runtime.R) {
	for _, service := range runtime.Services() {
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
