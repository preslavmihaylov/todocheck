package redmine

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Task model
type Task struct {
	Issue struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"issue"`
}

// GetStatus of redmine task, based on underlying structure
func (t *Task) GetStatus() taskstatus.TaskStatus {
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
