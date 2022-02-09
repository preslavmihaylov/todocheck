package jira

import "github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"

// Task JSON model as returned by the Jira Rest API
type Task struct {
	Fields struct {
		Status struct {
			StatusCategory struct {
				Name string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
	} `json:"fields"`
}

// GetStatus of jira task, based on underlying structure
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	switch t.Fields.Status.StatusCategory.Name {
	case "Done":
		return taskstatus.Closed, nil
	default:
		return taskstatus.Open, nil
	}
}
