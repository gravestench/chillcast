package twitch_integration

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/chillcast/pkg/services/config_file_manager"
)

// Ensure that Service implements the required interfaces.
var (
	_ servicemesh.Service                  = &Service{}
	_ servicemesh.HasLogger                = &Service{}
	_ servicemesh.HasDependencies          = &Service{}
	_ config_file_manager.HasDefaultConfig = &Service{}
)

// the following interfaces are used by the twitch integration service to
// discover and invoke the methods for binding event handlers.

// UsesEntireTwitchIntegration is an optional interface that represents a
// service capable of handling all Twitch integration events. By implementing
// this interface, the service can process events like connecting, whispers,
// private messages, chat clearing, room state, user notices, user state, global
// user state, notices, user joins/parts, reconnection, names list, pings, pongs,
// unset messages, and ping tracking. Implementing this interface provides
// comprehensive Twitch integration event coverage. Other services can implement
// individual handlers for specific events.
type UsesEntireTwitchIntegration interface {
	OnConnect
	OnWhisperMessage
	OnPrivateMessage
	OnClearChatMessage
	OnClearMessage
	OnRoomStateMessage
	OnUserNoticeMessage
	OnUserStateMessage
	OnGlobalUserStateMessage
	OnNoticeMessage
	OnUserJoinMessage
	OnUserPartMessage
	OnReconnectMessage
	OnNamesMessage
	OnPingMessage
	OnPongMessage
	OnUnsetMessage
	OnPingSent
}

// OnConnect is an interface that represents a handler for the OnConnect event,
// which is triggered when the service connects to Twitch.
type OnConnect interface {
	OnTwitchConnect()
}

// OnWhisperMessage is an interface that represents a handler for the OnWhisperMessage event,
// which is triggered when a whisper message is received from Twitch.
type OnWhisperMessage interface {
	OnTwitchWhisperMessage(message twitch.WhisperMessage)
}

// OnPrivateMessage is an interface that represents a handler for the OnPrivateMessage event,
// which is triggered when a private message is received from Twitch.
type OnPrivateMessage interface {
	OnTwitchPrivateMessage(message twitch.PrivateMessage)
}

// OnClearChatMessage is an interface that represents a handler for the OnClearChatMessage event,
// which is triggered when a clear chat message is received from Twitch.
type OnClearChatMessage interface {
	OnTwitchClearChatMessage(message twitch.ClearChatMessage)
}

// OnClearMessage is an interface that represents a handler for the OnClearMessage event,
// which is triggered when a clear message is received from Twitch.
type OnClearMessage interface {
	OnTwitchClearMessage(message twitch.ClearMessage)
}

// OnRoomStateMessage is an interface that represents a handler for the OnRoomStateMessage event,
// which is triggered when a room state message is received from Twitch.
type OnRoomStateMessage interface {
	OnTwitchRoomStateMessage(message twitch.RoomStateMessage)
}

// OnUserNoticeMessage is an interface that represents a handler for the OnUserNoticeMessage event,
// which is triggered when a user notice message is received from Twitch.
type OnUserNoticeMessage interface {
	OnTwitchUserNoticeMessage(message twitch.UserNoticeMessage)
}

// OnUserStateMessage is an interface that represents a handler for the OnUserStateMessage event,
// which is triggered when a user state message is received from Twitch.
type OnUserStateMessage interface {
	OnTwitchUserStateMessage(message twitch.UserStateMessage)
}

// OnGlobalUserStateMessage is an interface that represents a handler for the OnGlobalUserStateMessage event,
// which is triggered when a global user state message is received from Twitch.
type OnGlobalUserStateMessage interface {
	OnTwitchGlobalUserStateMessage(message twitch.GlobalUserStateMessage)
}

// OnNoticeMessage is an interface that represents a handler for the OnNoticeMessage event,
// which is triggered when a notice message is received from Twitch.
type OnNoticeMessage interface {
	OnTwitchNoticeMessage(message twitch.NoticeMessage)
}

// OnUserJoinMessage is an interface that represents a handler for the OnUserJoinMessage event,
// which is triggered when a user join message is received from Twitch.
type OnUserJoinMessage interface {
	OnTwitchUserJoinMessage(message twitch.UserJoinMessage)
}

// OnUserPartMessage is an interface that represents a handler for the OnUserPartMessage event,
// which is triggered when a user part message is received from Twitch.
type OnUserPartMessage interface {
	OnTwitchUserPartMessage(message twitch.UserPartMessage)
}

// OnReconnectMessage is an interface that represents a handler for the OnReconnectMessage event,
// which is triggered when a reconnect message is received from Twitch.
type OnReconnectMessage interface {
	OnTwitchReconnectMessage(message twitch.ReconnectMessage)
}

// OnNamesMessage is an interface that represents a handler for the OnNamesMessage event,
// which is triggered when a names message is received from Twitch.
type OnNamesMessage interface {
	OnTwitchNamesMessage(message twitch.NamesMessage)
}

// OnPingMessage is an interface that represents a handler for the OnPingMessage event,
// which is triggered when a ping message is received from Twitch.
type OnPingMessage interface {
	OnTwitchPingMessage(message twitch.PingMessage)
}

// OnPongMessage is an interface that represents a handler for the OnPongMessage event,
// which is triggered when a pong message is received from Twitch.
type OnPongMessage interface {
	OnTwitchPongMessage(message twitch.PongMessage)
}

// OnUnsetMessage is an interface that represents a handler for the OnUnsetMessage event,
// which is triggered when an unset message is received from Twitch.
type OnUnsetMessage interface {
	OnTwitchUnsetMessage(message twitch.RawMessage)
}

// OnPingSent is an interface that represents a handler for the OnPingSent event,
// which is triggered when a ping is sent to Twitch.
type OnPingSent interface {
	OnTwitchPingSent()
}
