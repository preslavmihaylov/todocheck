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

// TaskURLFrom taskID returns the url for the target github task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		return taskID[1:]
	}

	return taskID
}

// IssueAPIOrigin returns the URL for github's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	tokens := common.RemoveEmptyTokens(strings.Split(strings.ToLower(it.Origin), "/"))
	if !strings.HasPrefix(tokens[0], "http") {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, owner, repo := tokens[0], tokens[2], tokens[3]
	return fmt.Sprintf("%s//api.github.com/repos/%s/%s/issues/", scheme, owner, repo)
}
