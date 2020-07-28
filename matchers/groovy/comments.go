package groovy

import (
	"github.com/preslavmihaylov/todocheck/matchers/state"
)

// NewCommentMatcher for groovy comments
func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
	}
}

// CommentMatcher for groovy comments
type CommentMatcher struct {
	callback          state.CommentCallback
	buffer            string
	lines             []string
	linecnt           int
	stringToken       rune
	isMultiLineString bool
}

// NonCommentState for groovy comments
func (m *CommentMatcher) NonCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if currToken == '/' && nextToken == '/' {
		m.buffer += string(currToken)

		return state.SingleLineComment, nil
	} else if currToken == '/' && nextToken == '*' {
		m.buffer += string(currToken)
		m.lines = []string{line}
		m.linecnt = linecnt

		return state.MultiLineComment, nil
	} else if isSingleLineString(prevToken, currToken, nextToken) {
		m.stringToken = currToken

		return state.String, nil
	}

	return state.NonComment, nil
}

// StringState for groovy comments
func (m *CommentMatcher) StringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if m.isMultiLineString {
		return m.multiLineStringState(filename, line, linecnt, prevToken, currToken, nextToken)
	}

	return m.singleLineStringState(filename, line, linecnt, prevToken, currToken, nextToken)
}

func (m *CommentMatcher) singleLineStringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if !m.isMultiLineString && isMultiLineStringLiteral(m.stringToken, prevToken, currToken, nextToken) {
		m.isMultiLineString = true
		return state.String, nil
	}

	if prevToken != '\\' && currToken == m.stringToken {
		return state.NonComment, nil
	}

	return state.String, nil
}

func (m *CommentMatcher) multiLineStringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if isMultiLineStringLiteral(m.stringToken, prevToken, currToken, nextToken) {
		m.isMultiLineString = false

		// this is done to not consider the final string token as the beginning of a new string,
		// but rather, as the end of the current one.
		return state.String, nil
	}

	return state.String, nil
}

// SingleLineCommentState for groovy comments
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

// MultiLineCommentState for groovy comments
func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	m.buffer += string(currToken)
	if prevToken == '*' && currToken == '/' {
		err := m.callback(m.buffer, filename, m.lines, m.linecnt)
		if err != nil {
			return state.NonComment, err
		}

		m.resetState()
		return state.NonComment, nil
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
	m.isMultiLineString = false
}

func isSingleLineString(prevToken, currToken, nextToken rune) bool {
	return prevToken != '\\' && (currToken == '"' || currToken == '\'' || currToken == '`')
}

func isMultiLineStringLiteral(stringToken, prev, curr, next rune) bool {
	return prev == stringToken && curr == stringToken && next == stringToken
}
