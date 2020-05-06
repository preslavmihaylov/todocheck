package matchers

import (
	"errors"
)

// TodoMatcher for todo comments
type TodoMatcher interface {
	IsMatch(expr string) bool
	IsValid(expr string) bool
	ExtractIssueRef(expr string) (string, error)
}

// ErrInvalidTODO when passed todo expression is invalid
var ErrInvalidTODO = errors.New("invalid todo")
