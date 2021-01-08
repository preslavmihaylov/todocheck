package validation

import (
	"net/http"
	"testing"

	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

type mockIssueTracker struct{}

// returns a Task model, specific to the given issue tracker, which can be unmarshaled from JSON
func (m *mockIssueTracker) TaskModel() issuetracker.Task {
	panic("not implemented")
}

// IssueURLFor Returns the full URL for the issue
func (m *mockIssueTracker) IssueURLFor(taskID string) string {
	panic("not implemented")
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (m *mockIssueTracker) Exists() bool {
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (m *mockIssueTracker) InstrumentMiddleware(r *http.Request) error {
	panic("not implemented") // TODO: Implement
}

func TestInvalidOrigins(t *testing.T) {
	invalidConfigPaths := []string{
		"./fixtures/invalid/invalid_github_https.yaml",
		"./fixtures/invalid/invalid_github_origin.yaml",
		"./fixtures/invalid/invalid_github_www.yaml",
		"./fixtures/invalid/invalid_gitlab_origin.yaml",
		"./fixtures/invalid/invalid_gitlab_port.yaml",
		"./fixtures/invalid/invalid_issue_tracker.yaml",
		"./fixtures/invalid/invalid_jira_origin.yaml",
		"./fixtures/invalid/invalid_jira_port.yaml",
		"./fixtures/invalid/invalid_offline_url.yaml",
		"./fixtures/invalid/invalid_pivotal_origin.yaml",
		"./fixtures/invalid/invalid_redmine_origin.yaml",
		"./fixtures/invalid/invalid_redmine_port.yaml",
	}

	for _, path := range invalidConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := Validate(cfg, &mockIssueTracker{})
		if 0 == len(errors) {
			t.Errorf("%s should be invalid", path)
		}
	}
}

func TestValidOrigins(t *testing.T) {
	validConfigPaths := []string{
		"./fixtures/valid/valid_github_https.yaml",
		"./fixtures/valid/valid_github_origin.yaml",
		"./fixtures/valid/valid_github_www.yaml",
		"./fixtures/valid/valid_gitlab_origin.yaml",
		"./fixtures/valid/valid_gitlab_port.yaml",
		"./fixtures/valid/valid_gitlab_subdomain.yaml",
		"./fixtures/valid/valid_jira_origin.yaml",
		"./fixtures/valid/valid_jira_port.yaml",
		"./fixtures/valid/valid_jira_subdomain.yaml",
		"./fixtures/valid/valid_pivotal_origin.yaml",
		"./fixtures/valid/valid_redmine_origin.yaml",
		"./fixtures/valid/valid_redmine_port.yaml",
		"./fixtures/valid/valid_redmine_subdomain.yaml",
	}

	for _, path := range validConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := Validate(cfg, &mockIssueTracker{})
		if len(errors) > 0 {
			t.Errorf("%s should be valid but has errors: %v", path, errors)
		}
	}
}
