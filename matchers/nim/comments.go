package nim

import "github.com/preslavmihaylov/todocheck/matchers/state"

func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
		depth:    1,
	}
}

type CommentMatcher struct {
	callback    state.CommentCallback
	buffer      string
	lines       []string
	lineCount   int
	stringToken rune

	depth int
}

func (m *CommentMatcher) NonCommentState(
	filename,
	line string,
	lineCnt int,
	prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if isSingleLineOpener(currToken, nextToken) {
		// Catch single line comments
		m.buffer += string(currToken)
		return state.SingleLineComment, nil
	} else if isMultiLineOpener(currToken, nextToken) {
		// Catch multi line comments
		m.buffer += string(currToken)
		m.lines = []string{line}
		m.lineCount = lineCnt
		return state.MultiLineComment, nil
	} else if currToken == '"' || currToken == '\'' || currToken == '`' {
		m.stringToken = currToken
		return state.String, nil
	} else {
		m.buffer += string(currToken)
		return state.SingleLineComment, nil
	}
}

func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	m.buffer += string(currToken)

	if isMultiLineOpener(currToken, nextToken) {
		// Capture nested comments
		m.depth += 1
	} else if isMultiLineCloser(currToken, nextToken) {
		if m.depth > 1 {
			// Decrease nesting
			m.depth -= 1
		} else {
			// depth is at 1, this is a real comment closer
			err := m.callback(m.buffer, filename, m.lines, m.lineCount)
			if err != nil {
				return state.NonComment, err
			}

			m.buffer = ""
			m.lines = nil
			m.lineCount = 0
			m.stringToken = 0
			m.depth = 1

			return state.NonComment, nil
		}
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

		m = &CommentMatcher{
			callback:    m.callback,
			buffer:      "",
			lines:       nil,
			lineCount:   0,
			stringToken: 0,
			depth:       1,
		}
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
