package python

import (
	"fmt"

	"github.com/preslavmihaylov/todocheck/matchers/state"
)

// NewCommentMatcher for python comments
func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
	}
}

// CommentMatcher for python comments
type CommentMatcher struct {
	callback                  state.CommentCallback
	buffer                    string
	lines                     []string
	linecnt                   int
	stringToken               rune
	isExitingMultilineComment bool
}

// NonCommentState for python comments
func (m *CommentMatcher) NonCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if currToken == '#' && prevToken != '\\' {
		m.buffer += string(currToken)

		return state.SingleLineComment, nil
	} else if currToken == '"' || currToken == '\'' {
		m.stringToken = currToken

		return state.String, nil
	}

	return state.NonComment, nil
}

// StringState for python comments
func (m *CommentMatcher) StringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if isMultilineStringLiteral(m.stringToken, prevToken, currToken, nextToken) {
		m.buffer += string(prevToken) + string(currToken)
		m.lines = []string{line}
		m.linecnt = linecnt

		return state.MultiLineComment, nil
	} else if prevToken != '\\' && currToken == m.stringToken {
		return state.NonComment, nil
	}

	return state.String, nil
}

// SingleLineCommentState for python comments
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

// MultiLineCommentState for python comments
func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if m.isExitingMultilineComment {
		m.resetState()
		return state.NonComment, nil
	}

	m.buffer += string(currToken)
	if isMultilineStringLiteral(m.stringToken, prevToken, currToken, nextToken) {
		fmt.Println("exiting multi-line state")
		m.buffer += string(nextToken)
		err := m.callback(m.buffer, filename, m.lines, m.linecnt)
		if err != nil {
			return state.NonComment, err
		}

		m.isExitingMultilineComment = true
		return state.MultiLineComment, nil
	}

	if prevToken == '\n' {
		m.lines = append(m.lines, line)
	}

	return state.MultiLineComment, nil
}

func (m *CommentMatcher) resetState() {
	m.buffer = ""
	m.lines = nil
	m.linecnt = 0
	m.stringToken = 0
	m.isExitingMultilineComment = false
}

func isMultilineStringLiteral(stringToken, prev, curr, next rune) bool {
	return prev == stringToken && curr == stringToken && next == stringToken
}
