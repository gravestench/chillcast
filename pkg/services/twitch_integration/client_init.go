package twitch_integration

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"

	"golang.org/x/oauth2/clientcredentials"
	twitch2 "golang.org/x/oauth2/twitch"
)

func (s *Service) setupClient() {
	s.logger.Info("initializing twitch integration")

	cfg, err := s.Config()
	if err != nil {
		s.logger.Error("getting config", "error", err)
		panic(err)
	}

	creds := cfg.Group("credentials")
	userName := creds.GetString(keyUsername)
	oauthKey := creds.GetString(keyOauthKey)

	if !strings.HasPrefix(oauthKey, "oauth:") {
		oauthKey = "oauth:" + oauthKey
	}

	s.twitchIrcClient = twitch.NewClient(userName, oauthKey)

	s.twitchIrcClient.Join(userName)
	time.Sleep(time.Second)

	go func() {
		err = s.twitchIrcClient.Connect()
		if err != nil {
			s.logger.Warn("get your oauth token here: https://twitchapps.com/tmi/")

			cfgFilePath := filepath.Join(s.cfgManager.ConfigDirectory(), s.ConfigFileName())
			_ = s.cfgManager.SaveConfig(s.ConfigFileName())
			s.logger.Warn("edit your config file", "path", cfgFilePath)

			s.logger.Error("could not connect", "error", err)
			panic(err)
		}
	}()
}

func (s *Service) getToken(clientID, clientSecret string) string {
	oauth2Config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch2.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		s.logger.Error("getting access token", "error", err)
		panic(err)
	}

	return token.AccessToken
}
