package github

import (
	"fmt"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker implementation for integrating with public & private github issue trackers
type IssueTracker struct {
	Origin string
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
