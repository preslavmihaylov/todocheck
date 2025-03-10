package twig

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

// NewTodoMatcher for vue comments
func NewTodoMatcher(todos []string) *TodoMatcher {
	pattern := common.ArrayAsRegexAnyMatchExpression(todos)

	multiLineTodoPattern := regexp.MustCompile(`(?s)^\s*(<\!--|{#).*` + pattern)
	multiLineValidTodoPattern := regexp.MustCompile(`(?s)^\s*(<\!--|{#).*` + pattern + ` (#?[a-zA-Z0-9\-]+):.*`)

	return &TodoMatcher{
		multiLineTodoPattern:      multiLineTodoPattern,
		multiLineValidTodoPattern: multiLineValidTodoPattern,
	}
}

// TodoMatcher for vue comments
type TodoMatcher struct {
	multiLineTodoPattern      *regexp.Regexp
	multiLineValidTodoPattern *regexp.Regexp
}

// IsMatch checks if the current expression matches a vue comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return m.multiLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return m.multiLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	multiLineRes := m.multiLineValidTodoPattern.FindStringSubmatch(expr)
	if len(multiLineRes) >= 3 {
		return multiLineRes[2], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
