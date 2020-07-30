package models

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// PivotalTrackerTask model
type PivotalTrackerTask struct {
	CurrentState string `json:"current_state"`
}

// GetStatus of pivotal tracker task, based on underlying structure
func (t *PivotalTrackerTask) GetStatus() taskstatus.TaskStatus {
	switch t.CurrentState {
	case "finished":
		fallthrough
	case "delivered":
		fallthrough
	case "accepted":
		fallthrough
	case "rejected":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
