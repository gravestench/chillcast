package text_to_speech

import (
	"github.com/gravestench/runtime"
)

var (
	_ runtime.Service   = &Service{}
	_ runtime.HasLogger = &Service{}
)

type Dependency = ConvertsTextToSpeech

type ConvertsTextToSpeech interface {
	Speak(text string)
	SetVoice(voice string)
	Voices() []string
}
