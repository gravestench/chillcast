package twitch_integrated_tts

import (
	"log/slog"
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
	"github.com/gravestench/chillcast/pkg/services/profanity_detection"
	"github.com/gravestench/chillcast/pkg/services/text_to_speech"
)

// this service will just connect the TTS to the twitch integration service
type Service struct {
	logger *slog.Logger

	tts        text_to_speech.Dependency
	pd         profanity_detection.Dependency
	cfgManager config_file_manager.Dependency

	lastPersonThatSpoke string
	StartupTime         time.Time
	OnJoinDelay         time.Duration // prevent onJoin messages for a duration
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	// do nothing
}

func (s *Service) Name() string {
	return "Twitch-Integrated TTS"
}
