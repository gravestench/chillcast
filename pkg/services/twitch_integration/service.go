package twitch_integration

import (
	"log/slog"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

type Service struct {
	mesh       servicemesh.Mesh
	logger     *slog.Logger
	cfgManager config_file_manager.Manager

	twitchIrcClient *twitch.Client
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.mesh = mesh
	s.setupClient()
	s.bindService()
}

func (s *Service) Name() string {
	return "Twitch Integration"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
