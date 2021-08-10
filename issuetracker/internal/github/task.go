package github

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Task model
type Task struct {
	State string `json:"state"`
}

// GetStatus of github task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.State {
	case "closed":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
