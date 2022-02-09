package jira

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// New creates a new jira issuetracker instance
func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

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
	}

	switch it.AuthCfg.Type {
	case config.AuthTypeOffline:
		common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
		r.Header.Add("Authorization", "Bearer "+it.AuthCfg.Token)
	case config.AuthTypeAPIToken:
		common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
		data := []byte(fmt.Sprintf("%s:%s", it.AuthCfg.Options["username"], it.AuthCfg.Token))
		r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(data))
	default:
		return fmt.Errorf("unsupported authentication token type for jira: %s", it.AuthCfg.Type)
	}

	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for jira and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	switch it.AuthCfg.Type {
	case config.AuthTypeOffline:
		return fmt.Sprintf("Please go to %s and paste the offline token below.", it.AuthCfg.OfflineURL)
	case config.AuthTypeAPIToken:
		return "Please go to https://id.atlassian.com/manage-profile/security/api-tokens and paste the api token below."
	default:
		common.Assert(false, fmt.Sprintf(
			"token acquisition requested for unsupported authentication type: %s", it.AuthCfg.Type))
		return ""
	}
}

// TaskURLFrom taskID returns the url for the target Jira task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	return taskID
}

// IssueAPIOrigin returns the URL for Jira's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/rest/api/2/issue/", it.Origin)
}
