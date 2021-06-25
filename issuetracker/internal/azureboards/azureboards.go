package azureboards

import (
	"encoding/base64"
	"fmt"
	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"net/http"
	"strings"
)

func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{origin, authCfg}, nil
}

// IssueTracker implementation for integrating with public & private github issue trackers
type IssueTracker struct {
	Origin  string
	AuthCfg *config.Auth
}

func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

func (it *IssueTracker) Exists() bool {
	return true
}

// InstrumentMiddleware is a hook to instrument any necessary middleware for connecting with the issue tracker
func (it *IssueTracker) InstrumentMiddleware(r *http.Request) error {
	if it.AuthCfg == nil || it.AuthCfg.Type == config.AuthTypeNone {
		return nil
	} else if it.AuthCfg.Type != config.AuthTypeAPIToken {
		return fmt.Errorf("unsupported authentication token type for Azure Boards: %s", it.AuthCfg.Type)
	}

	common.Assert(it.AuthCfg.Token != "", "authentication token is empty")

	// Encode the token with Base64 and append ":". More on: https://bit.ly/35OZ4H8
	uEnc := base64.URLEncoding.EncodeToString([]byte(":" + it.AuthCfg.Token))

	r.Header.Add("Authorization", "Basic "+uEnc)
	return nil
}

// TokenAcquisitionInstructions returns instructions for manually acquiring the authentication token
// for github and the given authentication type
func (it *IssueTracker) TokenAcquisitionInstructions() string {
	if it.AuthCfg.Type == config.AuthTypeNone {
		return ""
	}

	return "Please create a read-only access token at Microsoft Azure & paste it here."
}

// taskURLFrom taskID returns the url for the target github task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		return taskID[1:]
	}
	// TODO Remove hardcoded API version. Make it configurable through yaml
	return fmt.Sprintf("/%s?api-version=6.0", taskID)
}

// issueAPIOrigin returns the URL for github's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/_apis/wit/workitems", it.repositoryURL())
}

func (it *IssueTracker) repositoryURL() string {
	scheme, owner, repo := it.urlTokensFromOrigin()
	return fmt.Sprintf("%s//dev.azure.com/%s/%s", scheme, owner, repo)

}

func (it *IssueTracker) urlTokensFromOrigin() (scheme, owner, repo string) {
	tokens := common.RemoveEmptyTokens(strings.Split(strings.ToLower(it.Origin), "/"))
	if !strings.HasPrefix(tokens[0], "http") {
		tokens = append([]string{"https:"}, tokens...)
	}
	scheme, owner, repo = tokens[0], tokens[2], tokens[3]
	return
}
