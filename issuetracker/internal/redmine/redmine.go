package redmine

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// New creates a new redmine issuetracker instance
func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

// IssueTracker implementation for integrating with public & private redmine issue trackers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

// TaskModel returns the model representing a deserialized redmine task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the redmine issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (it *IssueTracker) Exists() bool {
	// feature not supported for redmine yet
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeAPIToken {
		return fmt.Errorf("unsupported authentication token type for redmine: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
	r.Header.Add("X-Redmine-API-Key", it.AuthCfg.Token)
	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for pivotaltracker and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return ""
	}

	return fmt.Sprintf("Please go to %s/my/account, create a new API token & paste it here.", it.Origin)
}

// TaskURLFrom taskID returns the url for the target redmine task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		taskID = taskID[1:]
	}

	return taskID + ".json"
}

// IssueAPIOrigin returns the URL for redmine's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/issues/", it.Origin)
}
