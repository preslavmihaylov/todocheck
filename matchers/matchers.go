package matchers

import (
	"path/filepath"
	"sync"

	"github.com/preslavmihaylov/todocheck/matchers/groovy"
	"github.com/preslavmihaylov/todocheck/matchers/php"
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
	newTodoMatcher     func([]string) TodoMatcher
	newCommentsMatcher func(callback state.CommentCallback) CommentMatcher
}

var (
	once            sync.Once
	standardMatcher TodoMatcher
	scriptsMatcher  TodoMatcher
	phpMatcher      TodoMatcher
	pythonMatcher   TodoMatcher
	groovyMatcher   TodoMatcher
)

var (
	standardMatcherFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				if standardMatcher == nil {
					standardMatcher = standard.NewTodoMatcher(customTodos)
				}
			})
			return standardMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return standard.NewCommentMatcher(callback, false)
		},
	}
	standardMatcherWithNestedMultilineCommentsFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				if standardMatcher == nil {
					standardMatcher = standard.NewTodoMatcher(customTodos)
				}
			})
			return standardMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return standard.NewCommentMatcher(callback, true)
		},
	}
	scriptsMatcherFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				scriptsMatcher = scripts.NewTodoMatcher(customTodos)
			})
			return scriptsMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return scripts.NewCommentMatcher(callback)
		},
	}
	phpMatcherFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				phpMatcher = php.NewTodoMatcher(customTodos)
			})
			return phpMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return php.NewCommentMatcher(callback)
		},
	}
	pythonMatcherFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				pythonMatcher = python.NewTodoMatcher(customTodos)
			})
			return pythonMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return python.NewCommentMatcher(callback)
		},
	}
	groovyMatcherFactory = &matcherFactory{
		func(customTodos []string) TodoMatcher {
			once.Do(func() {
				groovyMatcher = groovy.NewTodoMatcher(customTodos)
			})
			return groovyMatcher
		},
		func(callback state.CommentCallback) CommentMatcher {
			return groovy.NewCommentMatcher(callback)
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
	".js":   standardMatcherFactory,
	".ts":   standardMatcherFactory,

	// file types, supporting standard comments \w nested multi-line comments
	".rs":    standardMatcherWithNestedMultilineCommentsFactory,
	".swift": standardMatcherWithNestedMultilineCommentsFactory,
	".scala": standardMatcherWithNestedMultilineCommentsFactory,
	".sc":    standardMatcherWithNestedMultilineCommentsFactory,

	// groovy file extensions
	".groovy": groovyMatcherFactory,
	".gvy":    groovyMatcherFactory,
	".gy":     groovyMatcherFactory,
	".gsh":    groovyMatcherFactory,

	// file types, supporting scripts comments
	".sh":   scriptsMatcherFactory,
	".bash": scriptsMatcherFactory,
	".zsh":  scriptsMatcherFactory,
	".R":    scriptsMatcherFactory,

	// file types, supporting php comments
	".php": phpMatcherFactory,

	// file types, supporting python comments
	".py": pythonMatcherFactory,
}

// TodoMatcherForFile gets the correct todo matcher for the given filename
func TodoMatcherForFile(filename string, todos []string) TodoMatcher {
	extension := filepath.Ext(filename)
	if matcherFactory, ok := supportedMatchers[extension]; ok {
		return matcherFactory.newTodoMatcher(todos)
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
