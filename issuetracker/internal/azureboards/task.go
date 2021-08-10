package azureboards

import (
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

// Task model
type Task struct {
	ID     int `json:"id"`
	Fields struct {
		State string `json:"System.State"`
	}
}

// GetStatus of github task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.Fields.State {
	case "Done":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
