package case_insensitive

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/matchers"
)

type TodoMatcher struct {
	matcher matchers.TodoMatcher
}

func NewTodoMatcher(matcher matchers.TodoMatcher) *TodoMatcher {
	return &TodoMatcher{
		matcher: matcher,
	}
}

func (m *TodoMatcher) IsMatch(expr string) bool {
	return m.matcher.IsMatch(todoToUpper(expr))
}

func (m *TodoMatcher) IsValid(expr string) bool {
	return m.matcher.IsValid(todoToUpper(expr))
}

func (m *TodoMatcher) ExtractIssueRef(expr string) (string, error) {
	return m.matcher.ExtractIssueRef(todoToUpper(expr))
}

func todoToUpper(expr string) string {
	re := regexp.MustCompile(`[Tt][Oo][Dd][Oo]`)
	return re.ReplaceAllString(expr, "TODO")
}
