package soundboard

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func (s *Service) playAudio(filePath string) {
	s.Logger().Info("playing audio file", "path", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		s.Logger().Error("opening audio file", "error", err)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		s.Logger().Error("decoding audio file", "error", err)
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
