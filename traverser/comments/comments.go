package comments

import (
	"strings"

	"github.com/preslavmihaylov/todocheck/logger"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/matchers/state"
	"github.com/preslavmihaylov/todocheck/traverser/lines"
)

// NewTraverser for comments
func NewTraverser(ignoredPaths []string, callback state.CommentCallback) *Traverser {
	return &Traverser{
		ignoredPaths:            ignoredPaths,
		supportedFileExtensions: matchers.SupportedFileExtensions(),
		state:                   state.NonComment,
		callback:                callback,
	}
}

// Traverser for comments in a given filename
type Traverser struct {
	ignoredPaths            []string
	supportedFileExtensions []string

	matcher  matchers.CommentMatcher
	filename string
	callback state.CommentCallback

	callbackErr error
	state       state.CommentState
}

// TraversePath and perform a callback on each line in each file
func (t *Traverser) TraversePath(path string) error {
	var prev, curr, next rune
	return lines.TraversePath(path, t.ignoredPaths, t.supportedFileExtensions, func(filename, line string, linecnt int) error {
		if !strings.HasSuffix(line, "\n") {
			line += "\n"
		}

		for _, b := range line {
			curr = next
			next = b
			err := t.handleStateChange(filename, line, linecnt, prev, curr, next)
			if err != nil {
				logger.Info("Error Occured during handleStateChange of file:" + filename + " at line: " + line)
			}
			prev = curr
		}

		curr = next
		next = 0
		err := t.handleStateChange(filename, line, linecnt, prev, curr, next)
		if err != nil {
			logger.Info("Error Occured during handleStateChange of file:" + filename + " at line: " + line)
		}

		prev = curr

		return t.callbackErr
	})
}

func (t *Traverser) handleStateChange(filename, line string, linecnt int, prevToken, currToken, nextToken rune) error {
	if t.callbackErr != nil {
		return t.callbackErr
	} else if filename != t.filename {
		t.matcher = matchers.CommentMatcherForFile(filename, t.callback)
		t.state = state.NonComment
		t.filename = filename

		// Our token traversal is actually one step behind the actual file,
		// so the very first time we start a new file, we need to skip the token
		return nil
	}

	var newState state.CommentState
	switch t.state {
	case state.NonComment:
		newState, t.callbackErr = t.matcher.NonCommentState(filename, line, linecnt, prevToken, currToken, nextToken)
	case state.String:
		newState, t.callbackErr = t.matcher.StringState(filename, line, linecnt, prevToken, currToken, nextToken)
	case state.SingleLineComment:
		newState, t.callbackErr = t.matcher.SingleLineCommentState(filename, line, linecnt, prevToken, currToken, nextToken)
	case state.MultiLineComment:
		newState, t.callbackErr = t.matcher.MultiLineCommentState(filename, line, linecnt, prevToken, currToken, nextToken)
	default:
		panic("unknown comment state")
	}

	t.state = newState

	return t.callbackErr
}
