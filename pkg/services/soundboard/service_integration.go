package soundboard

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
	"github.com/gravestench/chillcast/pkg/services/twitch_integration"
)

var (
	_ runtime.Service                      = &Service{}
	_ runtime.HasLogger                    = &Service{}
	_ runtime.HasDependencies              = &Service{}
	_ config_file_manager.HasDefaultConfig = &Service{}
	_ twitch_integration.OnPrivateMessage  = &Service{}
)

type Dependency = ConvertsTextToSpeech

type ConvertsTextToSpeech interface {
	Speak(text string)
	SetVoice(voice string)
	Voices() []string
}
