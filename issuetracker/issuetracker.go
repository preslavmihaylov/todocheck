package issuetracker

import (
	"errors"
	"fmt"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

// ErrUnsupportedHealthCheck is returned when the health check doesn't support the given issue tracker
var ErrUnsupportedHealthCheck = errors.New("unsupported issue tracker for health check")

// Task is an interface for generic task operations, decoupled from the specific platform's task structure
type Task interface {
	GetStatus() taskstatus.TaskStatus
}

// IssueTracker is an interface, which all issue tracker integration components adhere to in order to
// detach the specific issue trackers from the high-level rules for using issue trackers in the system
type IssueTracker interface {
	// returns a Task model, specific to the given issue tracker, which can be unmarshaled from JSON
	TaskModel() Task

	// TaskURLFrom the given task ID. Returns the appropriate task ID to append to rest api request.
	// In most use-cases, the taskID is returned as-is as it is simply appended to the issue tracker URL origin
	// (e.g. issuetracker.com/issues/{taskID}.
	//
	// However, some issue trackers might expect the taskID to be appended in a special format
	TaskURLFrom(taskID string) string

	// IssueAPIOrigin returns the task-fetching base url for the given issue tracker
	IssueAPIOrigin() string
}

// HealthcheckURL returns the health check base url given the issue tracker type and the site origin
func HealthcheckURL(issueTracker config.IssueTracker, origin string) (string, error) {
	switch issueTracker {
	case config.IssueTrackerGithub:
		scheme, owner, repo := parseGithubDetails(origin)
		return fmt.Sprintf("%s//api.github.com/repos/%s/%s", scheme, owner, repo), nil
	default:
		return "", ErrUnsupportedHealthCheck
	}
}

func parseGithubDetails(origin string) (scheme, owner, repo string) {
	tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
	if !strings.HasPrefix(tokens[0], "http") {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, owner, repo = tokens[0], tokens[2], tokens[3]
	return
}
