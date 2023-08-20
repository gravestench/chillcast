package soundboard

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func (s *Service) playAudio(filePath string) {
	s.Logger().Info().Msgf("playing audio file: %v", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		s.Logger().Error().Msgf("failed to open audio file: %w", err)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		s.Logger().Error().Msgf("failed to decode audio file: %w", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done // wait until the audio finishes playing
	return
}
