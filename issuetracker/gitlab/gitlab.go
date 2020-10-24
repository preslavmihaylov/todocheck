package gitlab

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/issuetracker"
)

// IssueTracker is an issue tracker implementation for integrating with public & private gitlab issue trackers
type IssueTracker struct {
	Origin string
}

// TaskModel returns the model representing a deserialized gitlab task
func (it *IssueTracker) TaskModel() issuetracker.Task {
	return &Task{}
}

// TaskURLFrom taskID returns the url for the target gitlab task ID to fetch
func (it *IssueTracker) TaskURLFrom(taskID string) string {
	return taskID
}

// IssueAPIOrigin returns the URL for github's issue-fetching API
func (it *IssueTracker) IssueAPIOrigin() string {
	tokens := common.RemoveEmptyTokens(strings.Split(it.Origin, "/"))
	if tokens[0] != "http:" && tokens[0] != "https:" {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, host, owner, repo := tokens[0], tokens[1], tokens[2], tokens[3]
	urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", owner, repo))
	return fmt.Sprintf("%s//%s/api/v4/projects/%s/issues/", scheme, host, urlEncodedProject)
}
