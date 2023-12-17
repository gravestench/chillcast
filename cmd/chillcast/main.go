package main

import (
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
	"github.com/gravestench/chillcast/pkg/services/profanity_detection"
	"github.com/gravestench/chillcast/pkg/services/soundboard"
	"github.com/gravestench/chillcast/pkg/services/text_to_speech"
	"github.com/gravestench/chillcast/pkg/services/twitch_integrated_tts"
	"github.com/gravestench/chillcast/pkg/services/twitch_integration"
)

func main() {
	app := servicemesh.New("Chillcast")

	app.Add(&config_file_manager.Service{RootDirectory: "~/.config/chillcast"})
	app.Add(&twitch_integration.Service{})
	app.Add(&text_to_speech.Service{})
	app.Add(&profanity_detection.Service{})
	app.Add(&soundboard.Service{})

	// this connects the twitch integration to the TTS
	app.Add(&twitch_integrated_tts.Service{
		StartupTime: time.Now(),
		OnJoinDelay: time.Second * 60,
	})

	app.Run()
}
