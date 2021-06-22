package azureboards

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Task model
type Task struct {
	State string `json:"state"`
}

// GetStatus of github task, based on underlying structure
func (t *Task) GetStatus() taskstatus.TaskStatus {
	switch t.State {
	case "closed":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
