package twitch_integrated_tts

import (
	"fmt"
	"strings"

	"github.com/hegedustibor/htgo-tts/voices"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

func (s *Service) ConfigFileName() string {
	return "twitch_integrated_text_to_speech.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	g := cfg.Group(s.Name())

	g.SetDefault("enable", true)
	g.SetDefault("remove profanity", false)
	g.SetDefault("announce user", true)
	g.SetDefault("system voice", voices.EnglishUK)
	g.SetDefault("user voice", voices.EnglishAU)

	g.Set("_help :: possible voices (language codes)", strings.Join([]string{
		voices.English,
		voices.EnglishUK,
		voices.EnglishAU,
		voices.Japanese,
		voices.German,
		voices.Spanish,
		voices.Russian,
		voices.Arabic,
		voices.Czech,
		voices.Danish,
		voices.Dutch,
		voices.Finnish,
		voices.Greek,
		voices.Hindi,
		voices.Hungarian,
		voices.Indonesian,
		voices.Khmer,
		voices.Latin,
		voices.Italian,
		voices.Norwegian,
		voices.Polish,
		voices.Slovak,
		voices.Swedish,
		voices.Thai,
		voices.Turkish,
		voices.Ukrainian,
		voices.Vietnamese,
		voices.Afrikaans,
		voices.Bulgarian,
		voices.Catalan,
		voices.Welsh,
		voices.Estonian,
		voices.French,
		voices.Gujarati,
		voices.Icelandic,
		voices.Javanese,
		voices.Kannada,
		voices.Korean,
		voices.Latvian,
		voices.Malayalam,
		voices.Marathi,
		voices.Malay,
		voices.Nepali,
		voices.Portuguese,
		voices.Romanian,
		voices.Sinhala,
		voices.Serbian,
		voices.Sundanese,
		voices.Tamil,
		voices.Telugu,
		voices.Tagalog,
		voices.Urdu,
		voices.Chinese,
		voices.Swahili,
		voices.Albanian,
		voices.Burmese,
		voices.Macedonian,
		voices.Armenian,
		voices.Croatian,
		voices.Esperanto,
		voices.Bosnian,
	}, ","))

	return
}
