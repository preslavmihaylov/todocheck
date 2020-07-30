package authmiddleware

import (
	"net/http"

	"github.com/preslavmihaylov/todocheck/config"
)

// Func used as callback to plug-in necessary auth headers
type Func func(r *http.Request)

// For creates a new auth middleware Func based on the given configuration
func For(cfg *config.Local) Func {
	if cfg.Auth == nil || cfg.Auth.Type == config.AuthTypeNone {
		return noAuthMiddleware()
	}

	assertInvariant(cfg.Auth.Token != "", "invariant violated. No token found for auth middleware")
	if cfg.Auth.Type == config.AuthTypeOffline {
		return authorizationTokenMiddleware(cfg.Auth.Token)
	} else if cfg.IssueTracker == config.IssueTrackerGithub && cfg.Auth.Type == config.AuthTypeAPIToken {
		return authorizationTokenMiddleware(cfg.Auth.Token)
	} else if cfg.IssueTracker == config.IssueTrackerGitlab && cfg.Auth.Type == config.AuthTypeAPIToken {
		return gitlabAPITokenMiddleware(cfg.Auth.Token)
	} else if cfg.IssueTracker == config.IssueTrackerPivotal && cfg.Auth.Type == config.AuthTypeAPIToken {
		return pivotalTrackerAPITokenMiddleware(cfg.Auth.Token)
	} else if cfg.IssueTracker == config.IssueTrackerRedmine && cfg.Auth.Type == config.AuthTypeAPIToken {
		return redmineAPITokenMiddleware(cfg.Auth.Token)
	}

	panic("couldn't derive authentication middleware based on the given local configuration")
}

func noAuthMiddleware() Func {
	return func(r *http.Request) {}
}

func authorizationTokenMiddleware(token string) Func {
	return func(r *http.Request) {
		r.Header.Add("Authorization", "Bearer "+token)
	}
}

func gitlabAPITokenMiddleware(token string) Func {
	return func(r *http.Request) {
		r.Header.Add("PRIVATE-TOKEN", token)
	}
}

func pivotalTrackerAPITokenMiddleware(token string) Func {
	return func(r *http.Request) {
		r.Header.Add("X-TrackerToken", token)
	}
}

func redmineAPITokenMiddleware(token string) Func {
	return func(r *http.Request) {
		r.Header.Add("X-Redmine-API-Key", token)
	}
}

func assertInvariant(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
