package pivotaltracker

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Task model
type Task struct {
	CurrentState string `json:"current_state"`
}

// GetStatus of pivotal tracker task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.CurrentState {
	case "finished":
		fallthrough
	case "delivered":
		fallthrough
	case "accepted":
		fallthrough
	case "rejected":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
