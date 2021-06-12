package youtrack

import (
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

const (
	TaskStateIndex = 2
	Value          = "value"
	IsResolved     = "isResolved"
)

type Task struct {
	CustomFields []map[string]interface{} `json:"customFields"`
}

// GetStatus of youtrack task, based on underlying structure
func (t *Task) GetStatus() taskstatus.TaskStatus {
	switch t.CustomFields[TaskStateIndex][Value].(map[string]interface{})[IsResolved].(bool) {
	case true:
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
