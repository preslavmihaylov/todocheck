package jira

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Status for Jira tasks
type Status struct {
	Name string `json:"name"`
}

// Fields for Jira tasks
type Fields struct {
	Status `json:"status"`
}

// Task JSON model as returned by the Jira Rest API
type Task struct {
	Fields `json:"fields"`
}

// GetStatus of jira task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.Fields.Status.Name {
	case "Done":
		fallthrough
	case "Closed":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
