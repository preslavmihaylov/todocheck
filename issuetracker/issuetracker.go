package issuetracker

import (
	"errors"

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

	// IssueURLFor Returns the full URL for the issue
	IssueURLFor(taskID string) string

	// Exists verifies if the issue tracker exists based on the provided configuration
	Exists() bool
}
