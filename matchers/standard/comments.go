package standard

import "github.com/preslavmihaylov/todocheck/matchers/state"

// NewCommentMatcher for standard comments
func NewCommentMatcher(callback state.CommentCallback, hasNestedMultilineComments bool) *CommentMatcher {
	return &CommentMatcher{
		callback:                   callback,
		hasNestedMultilineComments: hasNestedMultilineComments,
		currentDepth:               1,
	}
}

// CommentMatcher for standard comments
type CommentMatcher struct {
	callback    state.CommentCallback
	buffer      string
	lines       []string
	linecnt     int
	stringToken rune

	hasNestedMultilineComments bool
	currentDepth               int
}

// NonCommentState for standard comments
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
	m.buffer += string(currToken)
	if m.hasNestedMultilineComments && currToken == '/' && nextToken == '*' {
		m.currentDepth++
	} else if prevToken == '*' && currToken == '/' {
		if m.hasNestedMultilineComments && m.currentDepth > 1 {
			m.currentDepth--
		} else {
			err := m.callback(m.buffer, filename, m.lines, m.linecnt)
			if err != nil {
				return state.NonComment, err
			}

			m.resetState()
			return state.NonComment, nil
		}
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
	m.currentDepth = 1
}
