package github

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// New creates a new github issuetracker instance
func New(authCfg *config.Auth, origin string) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

// IssueTracker implementation for integrating with public & private github issue trackers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

// TaskModel returns the model representing a deserialized github task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the github issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (it *IssueTracker) Exists() bool {
	// in order to enable this feature for github, we first have to migrate the authMiddleware functionality from authmanager in the issuetracker interface
	// only then can we plug in the necessary apitoken to correctly check if the repository exists
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg == nil || it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeAPIToken {
		return fmt.Errorf("unsupported authentication token type for github: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")
	r.Header.Add("Authorization", "Bearer "+it.AuthCfg.Token)
	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for github and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return ""
	}

	return "Please go to https://github.com/settings/tokens, create a read-only access token & paste it here."
}

// taskURLFrom taskID returns the url for the target github task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		return taskID[1:]
	}

	return taskID
}

// issueAPIOrigin returns the URL for github's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/issues/", it.repositoryURL())
}

func (it *IssueTracker) repositoryURL() string {
	scheme, owner, repo := it.urlTokensFromOrigin()
	return fmt.Sprintf("%s//api.github.com/repos/%s/%s", scheme, owner, repo)
}

func (it *IssueTracker) urlTokensFromOrigin() (scheme, owner, repo string) {
	tokens := common.RemoveEmptyTokens(strings.Split(strings.ToLower(it.Origin), "/"))
	if !strings.HasPrefix(tokens[0], "http") {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, owner, repo = tokens[0], tokens[2], tokens[3]
	return
}
