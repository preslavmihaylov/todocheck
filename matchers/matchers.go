package matchers

import (
	"path/filepath"

	"github.com/preslavmihaylov/todocheck/matchers/python"
	"github.com/preslavmihaylov/todocheck/matchers/scripts"
	"github.com/preslavmihaylov/todocheck/matchers/standard"
	"github.com/preslavmihaylov/todocheck/matchers/state"
)

// TodoMatcher for todo comments
type TodoMatcher interface {
	IsMatch(expr string) bool
	IsValid(expr string) bool
	ExtractIssueRef(expr string) (string, error)
}

// CommentMatcher is used to match comments for various filetypes & comment-types.
// It is meant to be used by a file traversal state-machine
type CommentMatcher interface {
	NonCommentState(filename, line string, linecnt int, prevToken, currToken, nextToken rune) (state.CommentState, error)
	StringState(filename, line string, linecnt int, prevToken, currToken, nextToken rune) (state.CommentState, error)
	SingleLineCommentState(filename, line string, linecnt int, prevToken, currToken, nextToken rune) (state.CommentState, error)
	MultiLineCommentState(filename, line string, linecnt int, prevToken, currToken, nextToken rune) (state.CommentState, error)
}

type matcherFactory struct {
	newTodoMatcher     func() TodoMatcher
	newCommentsMatcher func(callback state.CommentCallback) CommentMatcher
}

var (
	standardMatcherFactory = &matcherFactory{
		func() TodoMatcher {
			return standard.NewTodoMatcher()
		},
		func(callback state.CommentCallback) CommentMatcher {
			return standard.NewCommentMatcher(callback)
		},
	}
	scriptsMatcherFactory = &matcherFactory{
		func() TodoMatcher {
			return scripts.NewTodoMatcher()
		},
		func(callback state.CommentCallback) CommentMatcher {
			return scripts.NewCommentMatcher(callback)
		},
	}
	pythonMatcherFactory = &matcherFactory{
		func() TodoMatcher {
			return python.NewTodoMatcher()
		},
		func(callback state.CommentCallback) CommentMatcher {
			return python.NewCommentMatcher(callback)
		},
	}
)

var supportedMatchers = map[string]*matcherFactory{
	// file types, supporting standard comments
	".go":   standardMatcherFactory,
	".java": standardMatcherFactory,
	".c":    standardMatcherFactory,
	".cpp":  standardMatcherFactory,
	".cs":   standardMatcherFactory,

	// file types, supporting scripts comments
	".sh":   scriptsMatcherFactory,
	".bash": scriptsMatcherFactory,
	".zsh":  scriptsMatcherFactory,

	// file types, supporting python comments
	".py": pythonMatcherFactory,
}

// TodoMatcherForFile gets the correct todo matcher for the given filename
func TodoMatcherForFile(filename string) TodoMatcher {
	extension := filepath.Ext(filename)
	if matcherFactory, ok := supportedMatchers[extension]; ok {
		return matcherFactory.newTodoMatcher()
	}

	return nil
}

// CommentMatcherForFile gets the correct comment matcher for the given filename
func CommentMatcherForFile(filename string, callback state.CommentCallback) CommentMatcher {
	extension := filepath.Ext(filename)
	if matcherFactory, ok := supportedMatchers[extension]; ok {
		return matcherFactory.newCommentsMatcher(callback)
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
