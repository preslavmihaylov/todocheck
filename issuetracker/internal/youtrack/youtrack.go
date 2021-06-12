package youtrack

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// New creates a new Youtrack IssueTracker instance
func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

// IssueTracker implementation for integrating with public & private Youtrack issue trackers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

// TaskModel returns the model representing a deserialized Youtrack task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the Youtrack issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// Exists verifies if the issue tracker exists based on the provided configuration
func (it *IssueTracker) Exists() bool {
	// feature not supported for Youtrack yet
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeAPIToken {
		return fmt.Errorf("unsupported authentication token type for youtrack: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")

	// add authorization header to the req
	r.Header.Add("Authorization", "Bearer "+it.AuthCfg.Token)
	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for Youtrack and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return ""
	}

	return fmt.Sprint(`Please go to https://www.jetbrains.com/help/youtrack/standalone/Manage-Permanent-Token.html,
					   follow the tutorial & paste the API token here.`)
}

// TaskURLFrom taskID returns the url for the target Youtrack task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		taskID = taskID[1:]
	}
	return fmt.Sprintf("%s?fields=customFields(value(isResolved))", taskID)
}

func (it *IssueTracker) urlTokensFromOrigin() (scheme, instance string) {
	tokens := common.RemoveEmptyTokens(strings.Split(strings.ToLower(it.Origin), "/"))
	if !strings.HasPrefix(tokens[0], "http") {
		tokens = append([]string{"https:"}, tokens...)
	}
	scheme, instance = tokens[0], tokens[1]
	return scheme, instance
}

func (it *IssueTracker) instanceURL() string {
	scheme, instance := it.urlTokensFromOrigin()
	return fmt.Sprintf("%s//%s", scheme, instance)
}

// IssueAPIOrigin returns the URL for Youtrack's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/youtrack/api/issues/", it.instanceURL())
}
