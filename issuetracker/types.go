package issuetracker

import "github.com/preslavmihaylov/todocheck/issuetracker/models"

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

// FromString converts a string-enoded issue tracker type to correct type
func FromString(str string) Type {
	switch str {
	case string(Jira):
		return Jira
	default:
		return Invalid
	}
}
