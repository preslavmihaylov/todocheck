package matchers

import (
	"errors"
	"path/filepath"
)

// TodoMatcher for todo comments
type TodoMatcher interface {
	IsMatch(expr string) bool
	IsValid(expr string) bool
	ExtractIssueRef(expr string) (string, error)
}

// ErrInvalidTODO when passed todo expression is invalid
var ErrInvalidTODO = errors.New("invalid todo")

// Supported file types
const (
	Go = ".go"
)

// ForFile gets the correct matcher for the given filename
func ForFile(filename string) TodoMatcher {
	switch filepath.Ext(filename) {
	case Go:
		return Standard()
	default:
		return nil
	}
}
