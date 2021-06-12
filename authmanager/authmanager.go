package authmanager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/preslavmihaylov/todocheck/authmanager/authstore"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	githubAPITokenMsg   = "Please go to https://github.com/settings/tokens, create a read-only access token & paste it here:\nToken: "
	gitlabAPITokenMsg   = "Please go to %s/profile/personal_access_tokens, create a read-only access token & paste it here:\nToken: "
	pivotalAPITokenMsg  = "Please go to https://www.pivotaltracker.com/profile, create a new API token & paste it here:\nToken: "
	redmineAPITokenMsg  = "Please go to %s/my/account, create a new API token & paste it here:\nToken: "
	youtrackAPITokenMsg = "Please go to https://www.jetbrains.com/help/youtrack/standalone/Manage-Permanent-Token.html, follow the tutorial, create a new API token & paste it here:\nToken: "

	authTokenEnvVariable = "TODOCHECK_AUTH_TOKEN"
)

// AcquireToken stores the issue tracker's auth token based on the auth type specified
func AcquireToken(cfg *config.Local, tracker issuetracker.IssueTracker) error {
	if !cfg.Auth.Type.IsValid() {
		return fmt.Errorf("invalid auth type: %q. valid auth types are: %q", cfg.Auth.Type, config.ValidAuthTypes)
	} else if cfg.Auth.Type == config.AuthTypeNone {
		return nil
	}

	tokenKey := cfg.Origin
	if cfg.Auth.Type == config.AuthTypeOffline {
		tokenKey = cfg.Auth.OfflineURL
	}

	instructions := tracker.TokenAcquisitionInstructions()
	if instructions == "" {
		panic("It's on us! We don't know how to handle this authentication token type." +
			" Please file an issue here - https://github.com/preslavmihaylov/todocheck/issues/new")
	}

	return acquireToken(cfg.Auth, tokenKey, instructions)
}

func acquireToken(authCfg *config.Auth, tokenKey string, instructions string) error {
	store, err := authstore.CreateIfNotExists(authCfg.TokensCache, authstore.DefaultConfigPermissions)
	if err != nil {
		return fmt.Errorf("couldn't read auth tokens config: %w", err)
	}

	if store.Tokens[tokenKey] != "" {
		authCfg.Token = store.Tokens[tokenKey]
		return nil
	} else if envToken := os.Getenv(authTokenEnvVariable); envToken != "" {
		authCfg.Token = envToken
		return nil
	}

	fmt.Printf("%s\nToken: ", instructions)
	tokenBs, err := readPassword()
	if err != nil {
		return fmt.Errorf("couldn't acquire token: %w", err)
	}

	return setAndPersistToken(authCfg, store, tokenKey, strings.TrimSpace(string(tokenBs)))
}

// Make token input scriptable, while preserving the hidden prompt behavior for users
// https://github.com/golang/go/issues/19909#issuecomment-399409958
func readPassword() ([]byte, error) {
	fd := int(syscall.Stdin)
	if terminal.IsTerminal(fd) {
		return terminal.ReadPassword(fd)
	}

	reader := bufio.NewReader(os.Stdin)
	return reader.ReadBytes('\n')
}

func setAndPersistToken(authCfg *config.Auth, store *authstore.Config, key, token string) error {
	authCfg.Token = token
	store.Tokens[key] = authCfg.Token
	return store.Save(authCfg.TokensCache)
}
