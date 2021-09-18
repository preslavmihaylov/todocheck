package nim

import (
	"github.com/preslavmihaylov/todocheck/matchers/state"
)

func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
		depth:    0,
	}
}

type CommentMatcher struct {
	callback    state.CommentCallback
	buffer      string
	lines       []string
	lineCount   int
	stringToken rune
	depth       int
}

func (m *CommentMatcher) NonCommentState(
	filename,
	line string,
	lineCount int,
	prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if isSingleLineOpener(currToken, nextToken) {
		m.buffer += string(currToken)
		return state.SingleLineComment, nil
	} else if isMultiLineOpener(currToken, nextToken) {
		m.buffer += string(currToken)
		m.lines = []string{line}
		m.lineCount = lineCount
		m.depth++
		return state.MultiLineComment, nil
	} else if currToken == '"' || currToken == '\'' || currToken == '`' {
		m.stringToken = currToken
		return state.String, nil
	} else {
		return state.NonComment, nil
	}
}

func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	m.buffer += string(currToken)
	if isMultiLineOpener(currToken, nextToken) {
		// Capture nested comments
		m.depth++
	} else if isMultiLineCloser(currToken, nextToken) {
		m.depth--
	}

	if m.depth == 0 {
		err := m.callback(m.buffer, filename, m.lines, m.lineCount)
		if err != nil {
			return state.NonComment, err
		}

		m.reset()
		return state.NonComment, nil
	}

	if prevToken == '\n' {
		m.lines = append(m.lines, line)
	}

	return state.MultiLineComment, nil
}

func (m *CommentMatcher) SingleLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if currToken == '\n' {
		// Reach end of line i.e. end of comment
		err := m.callback(m.buffer, filename, []string{line}, linecnt)
		if err != nil {
			return state.NonComment, err
		}

		m.reset()
		return state.NonComment, nil
	}

	m.buffer += string(currToken)
	return state.SingleLineComment, nil
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

func isSingleLineOpener(currToken, nextToken rune) bool {
	return currToken == '#' && nextToken != '['
}

func isMultiLineOpener(currToken, nextToken rune) bool {
	return currToken == '#' && nextToken == '['
}

func isMultiLineCloser(currToken, nextToken rune) bool {
	return currToken == ']' && nextToken == '#'
}

func (m *CommentMatcher) reset() {
	m.buffer = ""
	m.lines = nil
	m.lineCount = 0
	m.stringToken = 0
	m.depth = 0
}
