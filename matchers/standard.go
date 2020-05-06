package matchers

import "regexp"

var singleLineTodoPattern = regexp.MustCompile("^\\s*//\\s*TODO[ :]")
var singleLineValidTodoPattern = regexp.MustCompile("^\\s*// TODO ([a-zA-Z0-9\\-]+):.*")

// Standard comment matcher for java-like comments
func Standard() TodoMatcher { return &standardMatcher{} }

// Standard matcher for standard java-like comments
type standardMatcher struct{}

func (m *standardMatcher) IsMatch(expr string) bool {
	return singleLineTodoPattern.Match([]byte(expr))
}

func (m *standardMatcher) IsValid(expr string) bool {
	return singleLineValidTodoPattern.Match([]byte(expr))
}

func (m *standardMatcher) ExtractIssueRef(expr string) (string, error) {
	if !m.IsValid(expr) {
		return "", ErrInvalidTODO
	}

	return singleLineValidTodoPattern.FindStringSubmatch(expr)[1], nil
}
