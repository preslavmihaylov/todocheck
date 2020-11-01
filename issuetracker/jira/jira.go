package jira

import (
	"fmt"

	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker is an issue tracker implementation for integrating with private Jira servers
type IssueTracker struct {
	Origin string
}

// TaskModel returns the model representing a deserialized Jira task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// TaskURLFrom taskID returns the url for the target Jira task ID to fetch
func (it *IssueTracker) TaskURLFrom(taskID string) string {
	return taskID
}

// IssueAPIOrigin returns the URL for Jira's issue-fetching API
func (it *IssueTracker) IssueAPIOrigin() string {
	return fmt.Sprintf("%s/rest/api/latest/issue/", it.Origin)
}
