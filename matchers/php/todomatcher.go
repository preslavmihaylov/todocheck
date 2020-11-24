package php

import (
	"regexp"
	"sync"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

var lock sync.Mutex

var singleLineTodoPattern *regexp.Regexp
var singleLineValidTodoPattern *regexp.Regexp

var singleLineScriptTodoPattern *regexp.Regexp
var singleLineScriptValidTodoPattern *regexp.Regexp

var multiLineTodoPattern *regexp.Regexp
var multiLineValidTodoPattern *regexp.Regexp

// NewTodoMatcher for php comments
func NewTodoMatcher(todos []string) *TodoMatcher {
	lock.Lock()
	defer lock.Unlock()

	// Single line
	if singleLineTodoPattern == nil {
		singleLineTodoPattern = regexp.MustCompile("^\\s*//.*" + common.ArrayAsRegexAnyMatchExpression(todos))
	}
	if singleLineValidTodoPattern == nil {
		singleLineValidTodoPattern = regexp.MustCompile("^\\s*// " + common.ArrayAsRegexAnyMatchExpression(todos) + " (#?[a-zA-Z0-9\\-]+):.*")
	}

	// Script line
	if singleLineScriptTodoPattern == nil {
		singleLineScriptTodoPattern = regexp.MustCompile("^\\s*#.*" + common.ArrayAsRegexAnyMatchExpression(todos))
	}
	if singleLineScriptValidTodoPattern == nil {
		singleLineScriptValidTodoPattern = regexp.MustCompile("^\\s*# " + common.ArrayAsRegexAnyMatchExpression(todos) + " (#?[a-zA-Z0-9\\-]+):.*")
	}

	// Multiline line
	if multiLineTodoPattern == nil {
		multiLineTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*" + common.ArrayAsRegexAnyMatchExpression(todos))
	}
	if multiLineValidTodoPattern == nil {
		multiLineValidTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*" + common.ArrayAsRegexAnyMatchExpression(todos) + " (#?[a-zA-Z0-9\\-]+):.*")
	}

	return &TodoMatcher{
		singleLineTodoPattern:            singleLineTodoPattern,
		singleLineValidTodoPattern:       singleLineValidTodoPattern,
		singleLineScriptTodoPattern:      singleLineScriptTodoPattern,
		singleLineScriptValidTodoPattern: singleLineScriptValidTodoPattern,
		multiLineTodoPattern:             multiLineTodoPattern,
		multiLineValidTodoPattern:        multiLineValidTodoPattern,
	}
}

// TodoMatcher for php comments
type TodoMatcher struct {
	singleLineTodoPattern            *regexp.Regexp
	singleLineValidTodoPattern       *regexp.Regexp
	singleLineScriptTodoPattern      *regexp.Regexp
	singleLineScriptValidTodoPattern *regexp.Regexp
	multiLineTodoPattern             *regexp.Regexp
	multiLineValidTodoPattern        *regexp.Regexp
}

// IsMatch checks if the current expression matches a standard comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return m.singleLineTodoPattern.Match([]byte(expr)) ||
		m.singleLineScriptTodoPattern.Match([]byte(expr)) ||
		m.multiLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return m.singleLineValidTodoPattern.Match([]byte(expr)) ||
		m.singleLineScriptValidTodoPattern.Match([]byte(expr)) ||
		m.multiLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	singleLineRes := m.singleLineValidTodoPattern.FindStringSubmatch(expr)
	singleLineScriptRes := m.singleLineScriptValidTodoPattern.FindStringSubmatch(expr)
	multiLineRes := m.multiLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	} else if len(singleLineScriptRes) >= 2 {
		return singleLineScriptRes[1], nil
	} else if len(multiLineRes) >= 2 {
		return multiLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
