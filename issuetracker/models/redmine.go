package models

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// RedmineTask model
type RedmineTask struct {
	Issue struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"issue"`
}

// GetStatus of pivotal tracker task, based on underlying structure
func (t *RedmineTask) GetStatus() taskstatus.TaskStatus {
	switch t.Issue.Status.Name {
	case "Resolved":
		fallthrough
	case "Closed":
		fallthrough
	case "Feedback":
		fallthrough
	case "Rejected":
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
