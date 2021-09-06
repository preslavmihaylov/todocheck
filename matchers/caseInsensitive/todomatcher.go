package caseInsensitive

import (
	"regexp"

	"github.com/preslavmihaylov/todocheck/matchers"
)

type TodoMatcherCaseInsensitive struct {
	matcher matchers.TodoMatcher
}

func NewTodoMatcherCaseInsensitive(matcher matchers.TodoMatcher) *TodoMatcherCaseInsensitive {
	return &TodoMatcherCaseInsensitive{
		matcher: matcher,
	}
}

func (m *TodoMatcherCaseInsensitive) IsMatch(expr string) bool {
	return m.matcher.IsMatch(todoToUpper(expr))
}

func (m *TodoMatcherCaseInsensitive) IsValid(expr string) bool {
	return m.matcher.IsValid(todoToUpper(expr))
}

func (m *TodoMatcherCaseInsensitive) ExtractIssueRef(expr string) (string, error) {
	res, err := m.matcher.ExtractIssueRef(todoToUpper(expr))
	if err != nil {
		return "", err
	}

	return res, nil
}

func todoToUpper(expr string) string {
	re := regexp.MustCompile(`[Tt][Oo][Dd][Oo]`)
	return re.ReplaceAllString(expr, "TODO")
}
