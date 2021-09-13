package todoerrs

import (
	"fmt"

	"github.com/preslavmihaylov/todocheck/checker"
	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/matchers/state"
	"github.com/preslavmihaylov/todocheck/traverser/comments"
)

// TodoErrCallback is a function which acts on an encountered todo error
type TodoErrCallback func(todoerr *errors.TODO) error

// NewTraverser for todo errors
func NewTraverser(f *fetcher.Fetcher, ignoredPaths, customTodos []string, callback TodoErrCallback) *Traverser {
	return &Traverser{
		comments.NewTraverser(ignoredPaths, commentsCallback(checker.New(f), customTodos, callback)),
	}
}

// Traverser for todo errors
type Traverser struct {
	commentsTraverser *comments.Traverser
}

func commentsCallback(chk *checker.Checker, customTodos []string, todoErrCallback TodoErrCallback) state.CommentCallback {
	return func(comment, filepath string, lines []string, linecnt int) error {
		todoErr, err := chk.Check(matchers.TodoMatcherForFile(filepath, customTodos), comment, filepath, lines, linecnt)
		if err != nil {
			return fmt.Errorf("couldn't check todo line: %w", err)
		} else if todoErr != nil {
			err = todoErrCallback(todoErr)
			if err != nil {
				return fmt.Errorf("received error from todo err callback: %w", err)
			}
		}

		return nil
	}
}

// TraversePath for todo errors. Callback is invoked on encountered error
func (t *Traverser) TraversePath(path string) error {
	return t.commentsTraverser.TraversePath(path)
}
