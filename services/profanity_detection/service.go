package profanity_detection

import (
	"github.com/gravestench/runtime/pkg"
	"github.com/rs/zerolog"

	"github.com/TwiN/go-away"

	"github.com/gravestench/chillcast/services/config_file_manager"
)

type Service struct {
	logger *zerolog.Logger

	cfgManager config_file_manager.Dependency

	pd *goaway.ProfanityDetector
}

func (s *Service) Init(rt pkg.IsRuntime) {
	cfg, err := s.Config()
	if err != nil {
		s.Logger().Fatal().Msgf("getting config: %v", err)
	}

	g := cfg.Group(s.Name())

	profanities := g.GetStrings(cfgKeyProfanities)
	falsePositives := g.GetStrings(cfgKeyFalsePositives)
	falseNegatives := g.GetStrings(cfgKeyFalseNegatives)

	profanities = append(goaway.DefaultProfanities, profanities...)
	falsePositives = append(goaway.DefaultFalsePositives, falsePositives...)
	falseNegatives = append(goaway.DefaultFalseNegatives, falseNegatives...)

	s.pd = goaway.NewProfanityDetector().WithCustomDictionary(profanities, falsePositives, falseNegatives)
}

func (s *Service) Name() string {
	return "Profanity Detector"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) IsProfane(input string) bool {
	return s.pd.IsProfane(input)
}

func (s *Service) ExtractProfanity(input string) string {
	return s.pd.ExtractProfanity(input)
}

func (s *Service) Censor(input string) string {
	return s.pd.Censor(input)
}
