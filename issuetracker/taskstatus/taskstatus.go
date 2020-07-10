package taskstatus

// TaskStatus of linked issue
type TaskStatus int

// Supported task statuses
const (
	None TaskStatus = iota
	Open
	Closed
	NonExistent
)
