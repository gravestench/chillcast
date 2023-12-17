package text_to_speech

import (
	"github.com/gravestench/servicemesh"
)

var (
	_ servicemesh.Service   = &Service{}
	_ servicemesh.HasLogger = &Service{}
)

type Dependency = ConvertsTextToSpeech

type ConvertsTextToSpeech interface {
	Speak(text string)
	SetVoice(voice string)
	Voices() []string
}
