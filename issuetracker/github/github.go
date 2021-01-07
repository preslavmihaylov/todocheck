package github

import (
	"fmt"
	"net/http"
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
	res, err := http.Head(it.repositoryURL())
	if err != nil {
		return false
	} else if res.StatusCode == http.StatusNotFound {
		return false
	} else if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("received unexpected status code when making a request to the url - %s: %d",
			it.repositoryURL(), res.StatusCode))
	}

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
