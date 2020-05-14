package matchers

import (
	"regexp"
)

var singleLineTodoPattern = regexp.MustCompile("^\\s*//\\s*TODO[ :]")
var singleLineValidTodoPattern = regexp.MustCompile("^\\s*// TODO ([a-zA-Z0-9\\-]+):.*")

var multiLineTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*TODO")
var multiLineValidTodoPattern = regexp.MustCompile("(?s)^\\s*/\\*.*TODO ([a-zA-Z0-9\\-]+):.*")

// Standard comment matcher for java-like comments
func Standard() TodoMatcher { return &standardMatcher{} }

// Standard matcher for standard java-like comments
type standardMatcher struct{}

func (m *standardMatcher) IsMatch(expr string) bool {
	return singleLineTodoPattern.Match([]byte(expr)) || multiLineTodoPattern.Match([]byte(expr))
}

func (m *standardMatcher) IsValid(expr string) bool {
	return singleLineValidTodoPattern.Match([]byte(expr)) || multiLineValidTodoPattern.Match([]byte(expr))
}

func (m *standardMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", ErrInvalidTODO
	}

	singleLineRes := singleLineValidTodoPattern.FindStringSubmatch(expr)
	multiLineRes := multiLineValidTodoPattern.FindStringSubmatch(expr)
	if len(singleLineRes) >= 2 {
		return singleLineRes[1], nil
	} else if len(multiLineRes) >= 2 {
		return multiLineRes[1], nil
	}

	panic("Invariant violated. No issue reference found in valid TODO")
}
