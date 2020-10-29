package validation_test

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/validation"
	"github.com/preslavmihaylov/todocheck/config"
)

func TestInvalidOrigins(t *testing.T) {
	invalidConfigPaths := []string{
		"../testing/test_configs/invalid_github_https.yaml",
		"../testing/test_configs/invalid_github_origin.yaml",
		"../testing/test_configs/invalid_github_www.yaml",
		"../testing/test_configs/invalid_gitlab_origin.yaml",
		"../testing/test_configs/invalid_gitlab_port.yaml",
		"../testing/test_configs/invalid_issue_tracker.yaml",
		"../testing/test_configs/invalid_jira_origin.yaml",
		"../testing/test_configs/invalid_jira_port.yaml",
		"../testing/test_configs/invalid_offline_url.yaml",
		"../testing/test_configs/invalid_pivotal_origin.yaml",
		"../testing/test_configs/invalid_redmine_origin.yaml",
		"../testing/test_configs/invalid_redmine_port.yaml",
	}

	for _, path := range invalidConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := validation.Validate(cfg)
		if 0 == len(errors) {
			t.Errorf("%s should be invalid", path)
		}
	}
}

func TestValidOrigins(t *testing.T) {
	validConfigPaths := []string{
		"../testing/test_configs/valid_github_https.yaml",
		"../testing/test_configs/valid_github_origin.yaml",
		"../testing/test_configs/valid_github_www.yaml",
		"../testing/test_configs/valid_gitlab_origin.yaml",
		"../testing/test_configs/valid_gitlab_port.yaml",
		"../testing/test_configs/valid_gitlab_subdomain.yaml",
		"../testing/test_configs/valid_jira_origin.yaml",
		"../testing/test_configs/valid_jira_port.yaml",
		"../testing/test_configs/valid_jira_subdomain.yaml",
		"../testing/test_configs/valid_pivotal_origin.yaml",
		"../testing/test_configs/valid_redmine_origin.yaml",
		"../testing/test_configs/valid_redmine_port.yaml",
		"../testing/test_configs/valid_redmine_subdomain.yaml",
	}

	for _, path := range validConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := validation.Validate(cfg)
		if len(errors) > 0 {
			t.Errorf("%s should be valid but has errors: %v", path, errors)
		}
	}
}
