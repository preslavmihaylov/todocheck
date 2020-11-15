package groovy

import (
	"regexp"
	"sync"

	"github.com/preslavmihaylov/todocheck/matchers/errors"
)

var lock sync.Mutex

var singleLineTodoPattern *regexp.Regexp
var singleLineValidTodoPattern *regexp.Regexp

var multiLineTodoPattern *regexp.Regexp
var multiLineValidTodoPattern *regexp.Regexp

// NewTodoMatcher for groovy comments
func NewTodoMatcher(todos string) *TodoMatcher {
	lock.Lock()
	defer lock.Unlock()

	// Single line
	if singleLineTodoPattern == nil {
		singleLineTodoPattern = regexp.MustCompile("^\\s*//.*" + todos)
	}
	if singleLineValidTodoPattern == nil {
		singleLineValidTodoPattern = regexp.MustCompile("^\\s*// " + todos + " (#?[a-zA-Z0-9\\-]+):.*")
	}

	// Multiline line
	if multiLineTodoPattern == nil {
		multiLineTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*" + todos)
	}
	if multiLineValidTodoPattern == nil {
		multiLineValidTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*" + todos + " (#?[a-zA-Z0-9\\-]+):.*")
	}

	return &TodoMatcher{
		singleLineTodoPattern:      singleLineTodoPattern,
		singleLineValidTodoPattern: singleLineValidTodoPattern,
		multiLineTodoPattern:       multiLineTodoPattern,
		multiLineValidTodoPattern:  multiLineValidTodoPattern,
	}
}

// TodoMatcher for groovy comments
type TodoMatcher struct {
	singleLineTodoPattern      *regexp.Regexp
	singleLineValidTodoPattern *regexp.Regexp
	multiLineTodoPattern       *regexp.Regexp
	multiLineValidTodoPattern  *regexp.Regexp
}

// IsMatch checks if the current expression matches a groovy comment
func (m *TodoMatcher) IsMatch(expr string) bool {
	return m.singleLineTodoPattern.Match([]byte(expr)) || m.multiLineTodoPattern.Match([]byte(expr))
}

// IsValid checks if the expression is a valid todo comment
func (m *TodoMatcher) IsValid(expr string) bool {
	return m.singleLineValidTodoPattern.Match([]byte(expr)) || m.multiLineValidTodoPattern.Match([]byte(expr))
}

// ExtractIssueRef from the given expression.
// If the expression is invalid, an ErrInvalidTODO is returned
func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", errors.ErrInvalidTODO
	}

	singleLineRes := m.singleLineValidTodoPattern.FindStringSubmatch(expr)
	multiLineRes := m.multiLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	} else if len(multiLineRes) >= 2 {
		return multiLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
