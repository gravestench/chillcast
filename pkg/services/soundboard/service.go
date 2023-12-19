package soundboard

import (
	"log/slog"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

type Service struct {
	cfgManager config_file_manager.Dependency
	logger     *slog.Logger
	antispam   map[string]int
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.antispam = make(map[string]int)
}

func (s *Service) Name() string {
	return "Soundboard"
}

func (s *Service) Ready() bool {
	return true
}
