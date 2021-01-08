package config

import (
	"github.com/mitchellh/go-homedir"
)

// AuthType specifies the type of the auth token in todocheck's config
type AuthType string

// possible auth types
const (
	AuthTypeNone     AuthType = "none"
	AuthTypeOffline  AuthType = "offline"
	AuthTypeAPIToken AuthType = "apitoken"
)

//ValidAuthTypes is used for validation of auth type
var ValidAuthTypes = []AuthType{
	AuthTypeNone,
	AuthTypeOffline,
	AuthTypeAPIToken,
}

func defaultAuthCfg() *Auth {
	return &Auth{
		Type:        AuthTypeNone,
		TokensCache: DefaultTokensCache(),
	}
}

// Auth configuration section for specifying issue tracker auth options
type Auth struct {
	Type        AuthType `yaml:"type"`
	OfflineURL  string   `yaml:"offline_url"`
	TokensCache string   `yaml:"tokens_cache,omitempty"`
	Token       string   `yaml:"-"`
}

// DefaultTokensCache for storing auth tokens
func DefaultTokensCache() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic("couldn't read user home directory: " + err.Error())
	}

	return dir + "/.todocheck/authtokens.yaml"
}

// IsValid checks if the given AuthType is among the valid enum values
func (a AuthType) IsValid() bool {
	for _, auth := range ValidAuthTypes {
		if a == auth {
			return true
		}
	}

	return false
}

func (a AuthType) String() string {
	return string(a)
}
