// Package scripts contains a todo matcher & comments matcher for scripts.
// Scripts contain only single-line comments which begin with '#'
package scripts

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

var singleLineTodoPattern = regexp.MustCompile("^\\s*#.*TODO")
var singleLineValidTodoPattern = regexp.MustCompile("^\\s*# TODO ([a-zA-Z0-9\\-]+):.*")

// NewTodoMatcher for standard comments
func NewTodoMatcher() *TodoMatcher { return &TodoMatcher{} }

// TodoMatcher for standard comments
type TodoMatcher struct{}

// IsMatch checks if the current expression matches a standard comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return singleLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return singleLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	singleLineRes := singleLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
