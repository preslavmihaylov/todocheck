package models

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// GitlabTask model
type GitlabTask struct {
	State string `json:"state"`
}

// GetStatus of gitlab task, based on underlying structure
func (t *GitlabTask) GetStatus() taskstatus.TaskStatus {
	switch t.State {
	case "closed":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
