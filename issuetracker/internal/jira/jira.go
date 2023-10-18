package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

const defaultJIRAVersion = 9

// New creates a new jira issuetracker instance
func New(origin string, authCfg *config.Auth) (*IssueTracker, error) {
	return &IssueTracker{
		Origin:        origin,
		AuthCfg:       authCfg,
		serverVersion: deriveJIRAServerVersion(origin),
	}, nil
}

func deriveJIRAServerVersion(origin string) int {
	resp, err := http.Get(origin + "/rest/api/2/serverInfo")
	if err != nil {
		return defaultJIRAVersion
	}

	type serverInfo struct {
		Version string `json:"version"`
	}

	info := &serverInfo{}
	err = json.NewDecoder(resp.Body).Decode(info)
	if err != nil || info == nil || info.Version == "" {
		return defaultJIRAVersion
	}

	version, err := strconv.Atoi(strings.Split(info.Version, ".")[0])
	if err != nil {
		return defaultJIRAVersion
	}

	return version
}

// IssueTracker is an issue tracker implementation for integrating with private Jira servers
type IssueTracker struct {
	Origin        string
	AuthCfg       *config.Auth
	serverVersion int
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
		if it.serverVersion >= defaultJIRAVersion {
			data := []byte(fmt.Sprintf("%s:%s", it.AuthCfg.Options["username"], it.AuthCfg.Token))
			r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(data))
		} else {
			// versions prior to v8 use a Bearer authorization scheme
			r.Header.Add("Authorization", "Bearer "+it.AuthCfg.Token)
		}
	default:
		return fmt.Errorf("unsupported authentication token type for jira: %s", it.AuthCfg.Type)
	}

	return nil
}

// TaskURLFrom taskID returns the url for the target Jira task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	return taskID
}

// IssueAPIOrigin returns the URL for Jira's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/rest/api/2/issue/", it.Origin)
}
