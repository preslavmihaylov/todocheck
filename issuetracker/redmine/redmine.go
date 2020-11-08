package redmine

import (
	"fmt"

	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker implementation for integrating with public & private redmine issue trackers
type IssueTracker struct {
	Origin string
}

// TaskModel returns the model representing a deserialized redmine task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// IssueURLFor Returns the full URL for the redmine issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// TaskURLFrom taskID returns the url for the target redmine task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	return taskID + ".json"
}

// IssueAPIOrigin returns the URL for redmine's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
	return fmt.Sprintf("%s/issues/", it.Origin)
}
