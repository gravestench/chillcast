package text_to_speech

import (
	"time"

	"github.com/hegedustibor/htgo-tts/voices"
	"k8s.io/utils/strings/slices"
)

func (s *Service) Speak(text string) {
	s.logger.Info().Str("voice", s.speech.Language).Msg(text)

	if err := s.speech.Speak(text); err != nil {
		s.logger.Error().Msgf("could not convert text to speech: %v", err)
	}
}

func (s *Service) SetVoice(name string) {
	for s.speech.Folder == "" {
		time.Sleep(time.Second)
	}

	s.speech.Language = voices.English

	if slices.Contains(s.Voices(), name) {
		s.speech.Language = name
	}
}

func (s *Service) Voices() []string {
	return []string{
		voices.English, voices.EnglishUK, voices.EnglishAU, voices.Japanese,
		voices.German, voices.Spanish, voices.Russian, voices.Arabic,
		voices.Czech, voices.Danish, voices.Dutch, voices.Finnish, voices.Greek,
		voices.Hindi, voices.Hungarian, voices.Indonesian, voices.Khmer,
		voices.Latin, voices.Italian, voices.Norwegian, voices.Polish,
		voices.Slovak, voices.Swedish, voices.Thai, voices.Turkish,
		voices.Ukrainian, voices.Vietnamese, voices.Afrikaans, voices.Bulgarian,
		voices.Catalan, voices.Welsh, voices.Estonian, voices.French,
		voices.Gujarati, voices.Icelandic, voices.Javanese, voices.Kannada,
		voices.Korean, voices.Latvian, voices.Malayalam, voices.Marathi,
		voices.Malay, voices.Nepali, voices.Portuguese, voices.Romanian,
		voices.Sinhala, voices.Serbian, voices.Sundanese, voices.Tamil,
		voices.Telugu, voices.Tagalog, voices.Urdu, voices.Chinese,
		voices.Swahili, voices.Albanian, voices.Burmese, voices.Macedonian,
		voices.Armenian, voices.Croatian, voices.Esperanto, voices.Bosnian,
	}
}
