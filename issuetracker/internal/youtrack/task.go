package youtrack

import (
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

const (
	Value      = "value"
	Type       = "$type"
	StateType  = "StateIssueCustomField"
	IsResolved = "isResolved"
)

type Task struct {
	CustomFields []map[string]interface{} `json:"customFields"`
}

// GetStatus of youtrack task, based on underlying structure
func (t *Task) GetStatus() taskstatus.TaskStatus {
	isResolved := false
	// Loop through all fields to find status field
	for _, node := range t.CustomFields {
		if node[Type].(string) == StateType && node[Value] != nil {
			isResolved = node[Value].(map[string]interface{})[IsResolved].(bool)
			break
		}
	}

	switch isResolved {
	case true:
		return taskstatus.Closed
	default:
		return taskstatus.Open
	}
}
