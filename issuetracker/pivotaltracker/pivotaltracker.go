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

// IssueURLFor Returns the full URL for the pivotaltracker issue
func (it *IssueTracker) IssueURLFor(taskID string) string {
	return it.issueAPIOrigin() + it.taskURLFrom(taskID)
}

// TaskURLFrom taskID returns the url for the target pivotaltracker task ID to fetch
func (it *IssueTracker) taskURLFrom(taskID string) string {
	if strings.HasPrefix(taskID, "#") {
		return taskID[1:]
	}

	return taskID
}

// IssueAPIOrigin returns the URL for pivotaltracker's issue-fetching API
func (it *IssueTracker) issueAPIOrigin() string {
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
