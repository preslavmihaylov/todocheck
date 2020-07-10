package models

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Status for Jira tasks
type Status struct {
	Name string `json:"name"`
}

// Fields for Jira tasks
type Fields struct {
	Status `json:"status"`
}

// JiraTask JSON model as returned by Rest API
type JiraTask struct {
	Fields `json:"fields"`
}

// GetStatus of jira task, based on underlying structure
func (t *JiraTask) GetStatus() taskstatus.TaskStatus {
	switch t.Fields.Status.Name {
	case "Done":
		fallthrough
	case "Closed":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
