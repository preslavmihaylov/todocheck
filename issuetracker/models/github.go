package models

import "github.com/preslavmihaylov/todocheck/taskstatus"

// GithubTask model
type GithubTask struct {
	State string `json:"state"`
}

// GetStatus of jira task, based on underlying structure
func (t *GithubTask) GetStatus() taskstatus.TaskStatus {
	switch t.State {
	case "closed":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
