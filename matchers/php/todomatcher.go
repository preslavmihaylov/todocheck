package php

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

var singleLineTodoPattern = regexp.MustCompile("^\\s*//.*TODO")
var singleLineValidTodoPattern = regexp.MustCompile("^\\s*// TODO ([a-zA-Z0-9\\-]+):.*")

var singleLineScriptTodoPattern = regexp.MustCompile("^\\s*#.*TODO")
var singleLineScriptValidTodoPattern = regexp.MustCompile("^\\s*# TODO ([a-zA-Z0-9\\-]+):.*")

var multiLineTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*TODO")
var multiLineValidTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*TODO ([a-zA-Z0-9\\-]+):.*")

// NewTodoMatcher for php comments
func NewTodoMatcher() *TodoMatcher { return &TodoMatcher{} }

// TodoMatcher for php comments
type TodoMatcher struct{}

// IsMatch checks if the current expression matches a standard comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return singleLineTodoPattern.Match([]byte(expr)) ||
		singleLineScriptTodoPattern.Match([]byte(expr)) ||
		multiLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return singleLineValidTodoPattern.Match([]byte(expr)) ||
		singleLineScriptValidTodoPattern.Match([]byte(expr)) ||
		multiLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	singleLineRes := singleLineValidTodoPattern.FindStringSubmatch(expr)
	singleLineScriptRes := singleLineScriptValidTodoPattern.FindStringSubmatch(expr)
	multiLineRes := multiLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	} else if len(singleLineScriptRes) >= 2 {
		return singleLineScriptRes[1], nil
	} else if len(multiLineRes) >= 2 {
		return multiLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
