package main

import (
	"time"

	"github.com/gravestench/runtime"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
	"github.com/gravestench/chillcast/pkg/services/profanity_detection"
	"github.com/gravestench/chillcast/pkg/services/soundboard"
	"github.com/gravestench/chillcast/pkg/services/text_to_speech"
	"github.com/gravestench/chillcast/pkg/services/twitch_integrated_tts"
	"github.com/gravestench/chillcast/pkg/services/twitch_integration"
)

func main() {
	rt := runtime.New()

	rt.Add(&config_file_manager.Service{RootDirectory: "~/.config/chillcast"})
	rt.Add(&twitch_integration.Service{})
	rt.Add(&text_to_speech.Service{})
	rt.Add(&profanity_detection.Service{})
	rt.Add(&soundboard.Service{})

	// this connects the twitch integration to the TTS
	rt.Add(&twitch_integrated_tts.Service{
		StartupTime: time.Now(),
		OnJoinDelay: time.Second * 60,
	})

	rt.Run()
}
