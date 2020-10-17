package validation

import (
	"fmt"
	"net/url"

	"github.com/preslavmihaylov/todocheck/config"
)

// Validate validates the values of given configuration
func Validate(cfg *config.Local) []error {
	var errors []error

	if err := validateIssueTracker(cfg); err != nil {
		errors = append(errors, err)
	}

	if err := validateAuthOfflineURL(cfg); err != nil {
		errors = append(errors, err)
	}

	if err := validateIssueTrackerOrigin(cfg); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func validateIssueTracker(cfg *config.Local) error {
	if !cfg.IssueTracker.IsValid() {
		return fmt.Errorf("invalid issue tracker: %q is not supported", cfg.IssueTracker)
	}

	return nil
}

func validateAuthOfflineURL(cfg *config.Local) error {
	if _, err := url.ParseRequestURI(cfg.Auth.OfflineURL); cfg.Auth.Type == config.AuthTypeOffline && err != nil {
		return fmt.Errorf("invalid offline URL: %q", cfg.Auth.OfflineURL)
	}

	return nil
}

func validateIssueTrackerOrigin(cfg *config.Local) error {
	if cfg.IssueTracker != "" && !cfg.IssueTracker.IsValidOrigin(cfg.Origin) {
		return fmt.Errorf("%s is not a valid origin for issue tracker %s", cfg.Origin, cfg.IssueTracker)
	}

	return nil
}
