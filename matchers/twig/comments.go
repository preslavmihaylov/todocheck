package twig

import (
	"github.com/preslavmihaylov/todocheck/matchers/state"
)

// NewCommentMatcher for twig comments
func NewCommentMatcher(callback state.CommentCallback) *CommentMatcher {
	return &CommentMatcher{
		callback: callback,
	}
}

// CommentMatcher for twig comments
type CommentMatcher struct {
	callback                  state.CommentCallback
	buffer                    string
	lines                     []string
	linecnt                   int
	stringToken               rune
	isExitingMultilineComment bool
	commentType               string
	isStartingHTML            bool
}

// NonCommentState for twig comments
func (m *CommentMatcher) NonCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if prevToken == '{' && currToken == '#' {
		m.buffer += "{#"
		m.lines = []string{line}
		m.linecnt = linecnt
		m.commentType = "Twig"

		return state.MultiLineComment, nil
	} else if prevToken == '<' && currToken == '!' && nextToken == '-' {
		m.isStartingHTML = true

		return state.NonComment, nil
	} else if m.isStartingHTML && nextToken == '-' {
		m.buffer += "<!-"
		m.lines = []string{line}
		m.linecnt = linecnt
		m.commentType = "HTML"

		return state.MultiLineComment, nil
	} else if m.isStartingHTML && nextToken != '-' {
		m.isStartingHTML = false
	}

	return state.NonComment, nil
}

// StringState for twig comments
func (m *CommentMatcher) StringState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	if prevToken != '\\' && currToken == m.stringToken {
		return state.NonComment, nil
	}

	return state.String, nil
}

// SingleLineCommentState for twig comments
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

// MultiLineCommentState for twig comments
func (m *CommentMatcher) MultiLineCommentState(
	filename, line string, linecnt int, prevToken, currToken, nextToken rune,
) (state.CommentState, error) {
	m.buffer += string(currToken)

	if m.isExitingMultilineComment {

		err := m.callback(m.buffer, filename, m.lines, m.linecnt)
		if err != nil {
			return state.NonComment, err
		}

		m.resetState()

		return state.NonComment, nil
	}

	if isEndOfMultilineComment(m.commentType, prevToken, currToken, nextToken) {
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
	m.commentType = ""
	m.isStartingHTML = false
	m.isExitingMultilineComment = false
}

func isEndOfMultilineComment(commentType string, prev, curr, next rune) bool {
	if commentType == "Twig" {
		if curr == '#' && next == '}' {
			return true
		}
	}
	if commentType == "HTML" {
		if prev == '-' && curr == '-' && next == '>' {
			return true
		}
	}
	return false
}
