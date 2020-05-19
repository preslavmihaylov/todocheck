package config

import (
	"fmt"
	"strings"
	"syscall"

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
	fmt.Printf("Please go to %v and paste the offline token below:\n", a.OfflineURL)
	fmt.Print("Token: ")
	tokenBs, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("couldn't acquire offline token: %w", err)
	}

	a.Token = strings.TrimSpace(string(tokenBs))
	return nil
}
