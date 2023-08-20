package soundboard

import (
	"github.com/gravestench/runtime/pkg"
	"github.com/rs/zerolog"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

type Service struct {
	cfgManager config_file_manager.Dependency
	logger     *zerolog.Logger
	antispam   map[string]int
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) Init(rt pkg.IsRuntime) {
	s.antispam = make(map[string]int)
}

func (s *Service) Name() string {
	return "Soundboard"
}
