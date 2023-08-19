package twitch_integrated_tts

import (
	"time"

	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"github.com/gravestench/chillcast/services/config_file_manager"
	"github.com/gravestench/chillcast/services/profanity_detection"
	"github.com/gravestench/chillcast/services/text_to_speech"
)

// this service will just connect the TTS to the twitch integration service
type Service struct {
	logger *zerolog.Logger

	tts        text_to_speech.Dependency
	pd         profanity_detection.Dependency
	cfgManager config_file_manager.Dependency

	lastPersonThatSpoke string
	StartupTime         time.Time
	OnJoinDelay         time.Duration // prevent onJoin messages for a duration
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) Init(rt runtime.R) {
	// do nothing
}

func (s *Service) Name() string {
	return "Twitch-Integrated TTS"
}
