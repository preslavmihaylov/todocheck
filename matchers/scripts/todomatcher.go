package scripts

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

// NewTodoMatcher for scripts comments
func NewTodoMatcher(todos []string) *TodoMatcher {
	pattern := common.ArrayAsRegexAnyMatchExpression(todos)

	// Single line
	singleLineTodoPattern := regexp.MustCompile("^\\s*#.*" + pattern)
	singleLineValidTodoPattern := regexp.MustCompile("^\\s*# " + pattern + " (#?[a-zA-Z0-9\\-]+):.*")

	return &TodoMatcher{
		singleLineTodoPattern:      singleLineTodoPattern,
		singleLineValidTodoPattern: singleLineValidTodoPattern,
	}
}

// TodoMatcher for scripts comments
type TodoMatcher struct {
	singleLineTodoPattern      *regexp.Regexp
	singleLineValidTodoPattern *regexp.Regexp
}

// IsMatch checks if the current expression matches a standard comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return m.singleLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return m.singleLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	singleLineRes := m.singleLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
