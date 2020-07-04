package matchers

import (
	"path/filepath"

	"github.com/preslavmihaylov/todocheck/matchers/standard"
)

// TodoMatcher for todo comments
type TodoMatcher interface {
	IsMatch(expr string) bool
	IsValid(expr string) bool
	ExtractIssueRef(expr string) (string, error)
}

var (
	standardMatcher = standard.NewMatcher()
)

var supportedMatchers = map[string]TodoMatcher{
	".go":   standardMatcher,
	".java": standardMatcher,
	".c":    standardMatcher,
	".cpp":  standardMatcher,
	".cs":   standardMatcher,
}

// ForFile gets the correct matcher for the given filename
func ForFile(filename string) TodoMatcher {
	extension := filepath.Ext(filename)
	if matcher, ok := supportedMatchers[extension]; ok {
		return matcher
	}

	return nil
}

// SupportedFileExtensions for which there is a todo matcher
func SupportedFileExtensions() []string {
	var exts []string
	for ext := range supportedMatchers {
		exts = append(exts, ext)
	}

	return exts
}
