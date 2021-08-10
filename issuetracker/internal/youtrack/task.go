package youtrack

import (
	"errors"
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
func (t *Task) GetStatus() (taskstatus.TaskStatus, error) {
	isResolved := false
	var statusError error
	// Loop through all fields to find status field
	for _, node := range t.CustomFields {
		if node[Type].(string) == StateType {
			if node[Value] == nil {
				statusError = errors.New("couldn't fetch issue status from youtrack. This is probably on us, please file a bug report here: https://github.com/preslavmihaylov/todocheck/issues/new?title=failed%20get%20youtrack%20issue%20status&labels=bug")
			}
			isResolved = node[Value].(map[string]interface{})[IsResolved].(bool)
			break
		}
	}

	switch isResolved {
	case true:
		return taskstatus.Closed, statusError
	default:
		return taskstatus.Open, statusError
	}
}
