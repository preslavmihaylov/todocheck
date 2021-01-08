package jira

import (
	"fmt"
	"net/http"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker is an issue tracker implementation for integrating with private Jira servers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

// TaskModel returns the model representing a deserialized Jira task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the Jira issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (it *IssueTracker) Exists() bool {
	// feature not supported for Jira yet
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg == nil || it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeOffline {
		return fmt.Errorf("unsupported authentication token type for jira: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
	r.Header.Add("Authorization", "Bearer "+it.AuthCfg.Token)
	return nil
}

// TaskURLFrom taskID returns the url for the target Jira task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	return taskID
}

// IssueAPIOrigin returns the URL for Jira's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/rest/api/latest/issue/", it.Origin)
}
