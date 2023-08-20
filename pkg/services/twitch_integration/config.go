package twitch_integration

import (
	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

const (
	keyUsername = "Username"
	keyOauthKey = "Oauth Key"
)

func (s *Service) ConfigFileName() string {
	return "twitch_integration.json"
}

func (s *Service) Config() (*config_file_manager.Config, error) {
	return s.cfgManager.GetConfig(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config_file_manager.Config) {
	{
		g := cfg.Group("credentials")
		g.Set(keyUsername, "your username")
		g.Set(keyOauthKey, "your twitch oauth key")
	}

	{
		g := cfg.Group("handlers")

		g.Set("OnConnect", true)
		g.Set("OnWhisperMessage", true)
		g.Set("OnPrivateMessage", true)
		g.Set("OnClearChatMessage", true)
		g.Set("OnClearMessage", true)
		g.Set("OnRoomStateMessage", true)
		g.Set("OnUserNoticeMessage", true)
		g.Set("OnUserStateMessage", true)
		g.Set("OnGlobalUserStateMessage", true)
		g.Set("OnNoticeMessage", true)
		g.Set("OnUserJoinMessage", true)
		g.Set("OnUserPartMessage", true)
		g.Set("OnReconnectMessage", true)
		g.Set("OnNamesMessage", true)
		g.Set("OnPingMessage", true)
		g.Set("OnPongMessage", true)
		g.Set("OnUnsetMessage", true)
		g.Set("OnPingSent", true)
	}

	return
}
