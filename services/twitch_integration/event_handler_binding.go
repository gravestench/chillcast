package twitch_integration

import (
	"time"

	"github.com/gravestench/runtime"
)

// loopBindHandlers iterates over all services every second and binds event
// handlers for services that implement interfaces found in abstract.go
func (s *Service) loopBindHandlers(rt runtime.R) {
	// we will keep track service we've already bound
	bound := make(map[string]any)

	for {
		time.Sleep(time.Second * 1)

		for _, service := range rt.Services() {
			// if already bound, do nothing
			if _, isBound := bound[service.Name()]; isBound {
				continue
			}

			// otherwise bind, add to list
			s.bindService(service)
			bound[service.Name()] = service
		}
	}
}

func (s *Service) bindService(service runtime.Service) {
	if service == s {
		return
	}

	cfg, err := s.Config()
	if err != nil {
		s.Logger().Fatal().Msgf("getting config: %v", err)
	}

	cfgGroup := cfg.Group("handlers")

	if handler, ok := service.(OnConnect); ok {
		if cfgGroup.GetBool("OnConnect") {
			s.twitchIrcClient.OnConnect(handler.OnTwitchConnect)
		}
	}

	if handler, ok := service.(OnWhisperMessage); ok {
		if cfgGroup.GetBool("OnWhisperMessage") {
			s.twitchIrcClient.OnWhisperMessage(handler.OnTwitchWhisperMessage)
		}
	}

	if handler, ok := service.(OnPrivateMessage); ok {
		if cfgGroup.GetBool("OnPrivateMessage") {
			s.twitchIrcClient.OnPrivateMessage(handler.OnTwitchPrivateMessage)
		}
	}

	if handler, ok := service.(OnClearChatMessage); ok {
		if cfgGroup.GetBool("OnClearChatMessage") {
			s.twitchIrcClient.OnClearChatMessage(handler.OnTwitchClearChatMessage)
		}
	}

	if handler, ok := service.(OnClearMessage); ok {
		if cfgGroup.GetBool("OnClearMessage") {
			s.twitchIrcClient.OnClearMessage(handler.OnTwitchClearMessage)
		}
	}

	if handler, ok := service.(OnRoomStateMessage); ok {
		if cfgGroup.GetBool("OnRoomStateMessage") {
			s.twitchIrcClient.OnRoomStateMessage(handler.OnTwitchRoomStateMessage)
		}
	}

	if handler, ok := service.(OnUserNoticeMessage); ok {
		if cfgGroup.GetBool("OnUserNoticeMessage") {
			s.twitchIrcClient.OnUserNoticeMessage(handler.OnTwitchUserNoticeMessage)
		}
	}

	if handler, ok := service.(OnUserStateMessage); ok {
		if cfgGroup.GetBool("OnUserStateMessage") {
			s.twitchIrcClient.OnUserStateMessage(handler.OnTwitchUserStateMessage)
		}
	}

	if handler, ok := service.(OnGlobalUserStateMessage); ok {
		if cfgGroup.GetBool("OnGlobalUserStateMessage") {
			s.twitchIrcClient.OnGlobalUserStateMessage(handler.OnTwitchGlobalUserStateMessage)
		}
	}

	if handler, ok := service.(OnNoticeMessage); ok {
		if cfgGroup.GetBool("OnNoticeMessage") {
			s.twitchIrcClient.OnNoticeMessage(handler.OnTwitchNoticeMessage)
		}
	}

	if handler, ok := service.(OnUserJoinMessage); ok {
		if cfgGroup.GetBool("OnUserJoinMessage") {
			s.twitchIrcClient.OnUserJoinMessage(handler.OnTwitchUserJoinMessage)
		}
	}

	if handler, ok := service.(OnUserPartMessage); ok {
		if cfgGroup.GetBool("OnUserPartMessage") {
			s.twitchIrcClient.OnUserPartMessage(handler.OnTwitchUserPartMessage)
		}
	}

	if handler, ok := service.(OnReconnectMessage); ok {
		if cfgGroup.GetBool("OnReconnectMessage") {
			s.twitchIrcClient.OnReconnectMessage(handler.OnTwitchReconnectMessage)
		}
	}

	if handler, ok := service.(OnNamesMessage); ok {
		if cfgGroup.GetBool("OnNamesMessage") {
			s.twitchIrcClient.OnNamesMessage(handler.OnTwitchNamesMessage)
		}
	}

	if handler, ok := service.(OnPingMessage); ok {
		if cfgGroup.GetBool("OnPingMessage") {
			s.twitchIrcClient.OnPingMessage(handler.OnTwitchPingMessage)
		}
	}

	if handler, ok := service.(OnPongMessage); ok {
		if cfgGroup.GetBool("OnPongMessage") {
			s.twitchIrcClient.OnPongMessage(handler.OnTwitchPongMessage)
		}
	}

	if handler, ok := service.(OnUnsetMessage); ok {
		if cfgGroup.GetBool("OnUnsetMessage") {
			s.twitchIrcClient.OnUnsetMessage(handler.OnTwitchUnsetMessage)
		}
	}

	if handler, ok := service.(OnPingSent); ok {
		if cfgGroup.GetBool("OnPingSent") {
			s.twitchIrcClient.OnPingSent(handler.OnTwitchPingSent)
		}
	}
}
