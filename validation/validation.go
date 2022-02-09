package validation

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/fatih/color"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// Validate validates the values of given configuration
func Validate(cfg *config.Local, tracker issuetracker.IssueTracker) []error {
	var errs []error
	if err := validateIssueTracker(cfg); err != nil {
		errs = append(errs, err)
	}

	if cfg.Auth.Type == config.AuthTypeOffline {
		if err := validateAuthOfflineURLIsSet(cfg); err != nil {
			errs = append(errs, err)
		} else if err := validateAuthOfflineURL(cfg); err != nil {
			errs = append(errs, err)
		}
	}

	if err := validateIssueTrackerOrigin(cfg); err != nil {
		errs = append(errs, err)
	}

	if err := validateIssueTrackerExists(cfg, tracker); err != nil {
		errs = append(errs, err)
	}

	if err := validateIssueTrackerAuthType(cfg); err != nil {
		errs = append(errs, err)
	}

	if cfg.Auth.Token == "" && cfg.IssueTracker == config.IssueTrackerGithub {
		fmt.Fprintln(color.Output, color.YellowString(
			"WARNING: Github has API rate limits for all requests which do not contain a token.\n"+
				"         Please create a read-only access token to increase that limit.\n"+
				"         Go to https://developer.github.com/v3/#rate-limiting for more information."))
	}

	if cfg.IssueTracker == config.IssueTrackerJira && cfg.Auth.Type == config.AuthTypeAPIToken {
		if _, ok := cfg.Auth.Options["username"]; !ok {
			errs = append(errs, errors.New("api token authentication for JIRA requires username to be set - https://github.com/preslavmihaylov/todocheck#jira"))
		}
	}

	return errs
}

func validateIssueTracker(cfg *config.Local) error {
	if !cfg.IssueTracker.IsValid() {
		return fmt.Errorf("invalid issue tracker: %q is not supported", cfg.IssueTracker)
	}

	return nil
}

func validateAuthOfflineURL(cfg *config.Local) error {
	if _, err := url.ParseRequestURI(cfg.Auth.OfflineURL); err != nil {
		return fmt.Errorf("invalid offline URL: %q", cfg.Auth.OfflineURL)
	}

	return nil
}

func validateAuthOfflineURLIsSet(cfg *config.Local) error {
	if cfg.Auth.OfflineURL == "" {
		return fmt.Errorf("auth type chosen was %q but \"offline_url\" is not set", cfg.Auth.Type)
	}

	return nil
}

func validateIssueTrackerOrigin(cfg *config.Local) error {
	if cfg.IssueTracker != "" && !cfg.IssueTracker.IsValidOrigin(cfg.Origin) {
		return fmt.Errorf("%s is not a valid origin for issue tracker %s", cfg.Origin, cfg.IssueTracker)
	}

	return nil
}

func validateIssueTrackerExists(cfg *config.Local, tracker issuetracker.IssueTracker) error {
	if tracker.Exists() {
		return nil
	}

	if cfg.IssueTracker == config.IssueTrackerGithub {
		return fmt.Errorf("repository %s not found. Is the repository private? "+
			"More info: https://github.com/preslavmihaylov/todocheck#github", cfg.Origin)
	}

	return fmt.Errorf("repository %s not found", cfg.Origin)
}

func validateIssueTrackerAuthType(cfg *config.Local) error {
	if !cfg.IssueTracker.IsValidAuthType(cfg.Auth.Type) {
		return fmt.Errorf("unsupported authentication type for %s: %s", cfg.IssueTracker, cfg.Auth.Type.String())
	}
	return nil
}
