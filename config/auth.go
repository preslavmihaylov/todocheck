package config

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/preslavmihaylov/todocheck/config/authtokens"
	"golang.org/x/crypto/ssh/terminal"
)

// AuthType specifies the type of the auth token in todocheck's config
type AuthType string

// possible auth types
const (
	AuthTypeNone    AuthType = "none"
	AuthTypeOffline AuthType = "offline"
)

// Auth configuration section for specifying issue tracker auth options
type Auth struct {
	Type       AuthType `yaml:"type"`
	OfflineURL string   `yaml:"offline_url"`
	Token      string   `yaml:"-"`
}

// AcquireToken stores the issue tracker's auth token based on the auth type specified
func (a *Auth) AcquireToken() error {
	switch a.Type {
	case AuthTypeNone:
		return nil
	case AuthTypeOffline:
		return a.acquireOfflineToken()
	default:
		panic("unknown auth token type")
	}
}

func (a *Auth) acquireOfflineToken() error {
	tokensCfg, err := authtokens.CreateIfNotExists(authtokens.DefaultConfigFile(), authtokens.DefaultConfigPermissions)
	if err != nil {
		return fmt.Errorf("couldn't read auth tokens config: %w", err)
	}

	if tokensCfg.Tokens[a.OfflineURL] != "" {
		a.Token = tokensCfg.Tokens[a.OfflineURL]
		return nil
	}

	fmt.Printf("Please go to %v and paste the offline token below:\nToken: ", a.OfflineURL)
	tokenBs, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("couldn't acquire offline token: %w", err)
	}

	a.Token = strings.TrimSpace(string(tokenBs))
	tokensCfg.Tokens[a.OfflineURL] = a.Token
	tokensCfg.Save(authtokens.DefaultConfigFile())

	return nil
}
