package models

import "github.com/preslavmihaylov/todocheck/taskstatus"

// Task is an interface for generic task operations, decoupled from the specific platform's task structure
type Task interface {
	GetStatus() taskstatus.TaskStatus
}
