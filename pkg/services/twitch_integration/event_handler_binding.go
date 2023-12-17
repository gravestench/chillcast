package twitch_integration

import (
	"github.com/gempir/go-twitch-irc/v2"
)

func (s *Service) bindService() {
	cfg, err := s.Config()
	if err != nil {
		s.Logger().Error("getting config", "error", err)
		panic(err)
	}

	cfgGroup := cfg.Group("handlers")

	s.twitchIrcClient.OnConnect(func() {
		if !cfgGroup.GetBool("OnConnect") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnConnect); ok {
				handler.OnTwitchConnect()
			}
		}
	})

	s.twitchIrcClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		if !cfgGroup.GetBool("OnWhisperMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnWhisperMessage); ok {
				handler.OnTwitchWhisperMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if !cfgGroup.GetBool("OnPrivateMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnPrivateMessage); ok {
				handler.OnTwitchPrivateMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		if !cfgGroup.GetBool("OnClearChatMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnClearChatMessage); ok {
				handler.OnTwitchClearChatMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnClearMessage(func(message twitch.ClearMessage) {
		if !cfgGroup.GetBool("OnClearMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnClearMessage); ok {
				handler.OnTwitchClearMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnRoomStateMessage(func(message twitch.RoomStateMessage) {
		if !cfgGroup.GetBool("OnRoomStateMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnRoomStateMessage); ok {
				handler.OnTwitchRoomStateMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		if !cfgGroup.GetBool("OnUserNoticeMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnUserNoticeMessage); ok {
				handler.OnTwitchUserNoticeMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnUserStateMessage(func(message twitch.UserStateMessage) {
		if !cfgGroup.GetBool("OnUserStateMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnUserStateMessage); ok {
				handler.OnTwitchUserStateMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnGlobalUserStateMessage(func(message twitch.GlobalUserStateMessage) {
		if !cfgGroup.GetBool("OnGlobalUserStateMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnGlobalUserStateMessage); ok {
				handler.OnTwitchGlobalUserStateMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnNoticeMessage(func(message twitch.NoticeMessage) {
		if !cfgGroup.GetBool("OnNoticeMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnNoticeMessage); ok {
				handler.OnTwitchNoticeMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnUserJoinMessage(func(message twitch.UserJoinMessage) {
		if !cfgGroup.GetBool("OnUserJoinMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnUserJoinMessage); ok {
				handler.OnTwitchUserJoinMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnUserPartMessage(func(message twitch.UserPartMessage) {
		if !cfgGroup.GetBool("OnUserPartMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnUserPartMessage); ok {
				handler.OnTwitchUserPartMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnReconnectMessage(func(message twitch.ReconnectMessage) {
		if !cfgGroup.GetBool("OnReconnectMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnReconnectMessage); ok {
				handler.OnTwitchReconnectMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnNamesMessage(func(message twitch.NamesMessage) {
		if !cfgGroup.GetBool("OnNamesMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnNamesMessage); ok {
				handler.OnTwitchNamesMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnPingMessage(func(message twitch.PingMessage) {
		if !cfgGroup.GetBool("OnPingMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnPingMessage); ok {
				handler.OnTwitchPingMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnPongMessage(func(message twitch.PongMessage) {
		if !cfgGroup.GetBool("OnPongMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnPongMessage); ok {
				handler.OnTwitchPongMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnUnsetMessage(func(message twitch.RawMessage) {
		if !cfgGroup.GetBool("OnUnsetMessage") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnUnsetMessage); ok {
				handler.OnTwitchUnsetMessage(message)
			}
		}

	})

	s.twitchIrcClient.OnPingSent(func() {
		if !cfgGroup.GetBool("OnPingSent") {
			return
		}

		for _, service := range s.mesh.Services() {
			if handler, ok := service.(OnPingSent); ok {
				handler.OnTwitchPingSent()
			}
		}
	})
}
