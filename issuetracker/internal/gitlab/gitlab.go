package gitlab

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// New creates a new gitlab issuetracker instance
func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

// IssueTracker is an issue tracker implementation for integrating with public & private gitlab issue trackers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

// TaskModel returns the model representing a deserialized gitlab task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the gitlab issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (it *IssueTracker) Exists() bool {
	// feature not supported for gitlab yet
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg == nil || it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeAPIToken {
		return fmt.Errorf("unsupported authentication token type for gitlab: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
	r.Header.Add("PRIVATE-TOKEN", it.AuthCfg.Token)
	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for gitlab and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return ""
	}

	return fmt.Sprintf("Please go to %s/profile/personal_access_tokens, "+
		"create a read-only access token & paste it here.", extractBaseURL(it.Origin))
}

// TaskURLFrom taskID returns the url for the target gitlab task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		return taskID[1:]
	}

	return taskID
}

// IssueAPIOrigin returns the URL for github's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	tokens := common.RemoveEmptyTokens(strings.Split(strings.ToLower(it.Origin), "/"))
	if !strings.HasPrefix(tokens[0], "http:") && !strings.HasPrefix(tokens[0], "https:") {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, host, owner, repo := tokens[0], tokens[1], tokens[2], tokens[3]
	urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", owner, repo))
	return fmt.Sprintf("%s//%s/api/v4/projects/%s/issues/", scheme, host, urlEncodedProject)
}

func extractBaseURL(origin string) string {
	tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
	if tokens[0] != "http:" && tokens[0] != "https:" {
		return fmt.Sprintf("https://%s", tokens[0])
	}

	return fmt.Sprintf("%s//%s", tokens[0], tokens[1])
}
