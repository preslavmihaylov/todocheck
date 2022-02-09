package jira

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

type StatusCategory struct {
	Name string `json:"name"`
}

type Status struct {
	StatusCategory StatusCategory `json:"statusCategory"`
}

type Fields struct {
	Status Status `json:"status"`
}

// Task JSON model as returned by the Jira Rest API
type Task struct {
	Fields Fields `json:"fields"`
}

// GetStatus of jira task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.Fields.Status.StatusCategory.Name {
	case "Done":
		fallthrough
	case "Closed":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
