package issuetracker

import (
	"errors"

	"github.com/preslavmihaylov/todocheck/issuetracker/models"
)

// Type of issue tracker enum
type Type string

// Issue tracker types
const (
	Invalid Type = ""
	Jira         = "JIRA"
)

// TaskFor gets the corresponding task model, based on the issue tracker type
func TaskFor(issueTracker Type) models.Task {
	switch issueTracker {
	case Jira:
		return &models.JiraTask{}
	default:
		return nil
	}
}

// BaseURLFor returns the task-fetching base url given the issue tracker type and the site origin
func BaseURLFor(issueTracker Type, origin string) (string, error) {
	switch issueTracker {
	case Jira:
		return origin + "/rest/api/latest/issue/", nil
	default:
		return "", errors.New("unknown issue tracker type")
	}
}

// FromString converts a string-enoded issue tracker type to correct type
func FromString(str string) Type {
	switch str {
	case Jira:
		return Jira
	default:
		return Invalid
	}
}
