package pivotaltracker

import (
	"fmt"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker implementation for integrating with public pivotaltracker issue trackers
type IssueTracker struct {
	Origin string
}

// TaskModel returns the model representing a deserialized pivotaltracker task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// TaskURLFrom taskID returns the url for the target pivotaltracker task ID to fetch
func (it *IssueTracker) TaskURLFrom(taskID string) string {
	if len(taskID) > 0 && taskID[0] == '#' {
		return taskID[1:]
	}

	return taskID
}

// IssueAPIOrigin returns the URL for pivotaltracker's issue-fetching API
func (it *IssueTracker) IssueAPIOrigin() string {
	tokens := common.RemoveEmptyTokens(strings.Split(it.Origin, "/"))
	if tokens[0] == "pivotaltracker.com" {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, project := tokens[0], tokens[3]
	if tokens[2] == "n" {
		project = tokens[4]
	}

	return fmt.Sprintf("%s//www.pivotaltracker.com/services/v5/projects/%s/stories/", scheme, project)
}
