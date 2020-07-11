package scripts

import (
	"errors"

	"github.com/preslavmihaylov/todocheck/matchers/state"
)

// NewCommentMatcher for standard comments
func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
	}
}

// CommentMatcher for standard comments
type CommentMatcher struct {
	callback    state.CommentCallback
	buffer      string
	lines       []string
	linecnt     int
	stringToken rune
}

// NonCommentState for standard comments
func (m *CommentMatcher) NonCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if currToken == '#' {
		m.buffer += string(currToken)

		return state.SingleLineComment, nil
	} else if currToken == '"' || currToken == '\'' || currToken == '`' {
		m.stringToken = currToken

		return state.String, nil
	}

	return state.NonComment, nil
}

// StringState for standard comments
func (m *CommentMatcher) StringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if prevToken != '\\' && currToken == m.stringToken {
		return state.NonComment, nil
	}

	return state.String, nil
}

// SingleLineCommentState for standard comments
func (m *CommentMatcher) SingleLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if currToken == '\n' {
		err := m.callback(m.buffer, filename, []string{line}, linecnt)
		if err != nil {
			return state.NonComment, err
		}

		m.resetState()
		return state.NonComment, nil
	}

	m.buffer += string(currToken)
	return state.SingleLineComment, nil
}

// MultiLineCommentState for standard comments
func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	return state.NonComment,
		errors.New("invariant violated. We're in a multiline comment state, but only single-line comments exist for scripts")
}

func (m *CommentMatcher) resetState() {
	m.buffer = ""
	m.lines = nil
	m.linecnt = 0
	m.stringToken = 0
}
