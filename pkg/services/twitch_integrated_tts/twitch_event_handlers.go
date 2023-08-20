package twitch_integrated_tts

import (
	"fmt"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/hegedustibor/htgo-tts/voices"
)

func (s *Service) OnTwitchPrivateMessage(message twitch.PrivateMessage) {
	cfg, err := s.Config()
	if err != nil {
		s.Logger().Fatal().Msgf("getting config: %v", err)
	}

	g := cfg.Group(s.Name())

	if g.GetBool("remove profanity") && s.pd.IsProfane(message.Message) {
		return
	}

	if strings.Contains(message.Message, "http") {
		return
	}

	if s.lastPersonThatSpoke != message.User.Name {
		s.lastPersonThatSpoke = message.User.Name

		name := strings.ReplaceAll(message.User.Name, "_", "")
		s.tts.SetVoice(voices.EnglishUK)
		s.tts.Speak(name + " says: ")
	}

	s.tts.SetVoice(voices.EnglishAU)
	s.tts.Speak(message.Message)
}

func (s *Service) OnTwitchUserJoinMessage(message twitch.UserJoinMessage) {
	if time.Since(s.StartupTime) < s.OnJoinDelay {
		return
	}

	s.tts.SetVoice(voices.EnglishAU)
	s.tts.Speak(fmt.Sprintf("user %s has joined the chat", message.User))
}
