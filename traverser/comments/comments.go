package comments

import (
	"github.com/preslavmihaylov/todocheck/traverser/comments/state"
	"github.com/preslavmihaylov/todocheck/traverser/lines"
)

// CommentCallback is a function which acts on an encountered comment
type CommentCallback func(comment, filepath string, lines []string, linecnt int) error

// New comment traverser initialization
func New(callback CommentCallback) *Traverser {
	return &Traverser{
		state:    state.NonComment,
		callback: callback,
	}
}

// Traverser for comments in a given filepath
type Traverser struct {
	callback CommentCallback
	state    state.CommentState

	stringToken rune
	buffer      string
	filepath    string
	lines       []string
	linecnt     int
}

// TraversePath and perform a callback on each line in each file
func (t *Traverser) TraversePath(path string) error {
	var prev, curr, next rune
	return lines.TraversePath(path, func(filepath, line string, linecnt int) error {
		for _, b := range line {
			curr = next
			next = b
			t.handleStateChange(filepath, line, linecnt, prev, curr, next)

			prev = curr
		}

		curr = next
		next = 0
		t.handleStateChange(filepath, line, linecnt, prev, curr, next)

		prev = curr

		return nil
	})
}

func (t *Traverser) handleStateChange(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) {
	if t.filepath != "" && filepath != t.filepath {
		t.resetState()
		t.state = state.NonComment

		return
	}

	var newState state.CommentState
	switch t.state {
	case state.NonComment:
		newState = t.nonCommentState(filepath, line, linecnt, prevToken, currToken, nextToken)
	case state.String:
		newState = t.stringState(filepath, line, linecnt, prevToken, currToken, nextToken)
	case state.SingleLineComment:
		newState = t.singleLineCommentState(filepath, line, linecnt, prevToken, currToken, nextToken)
	case state.MultiLineComment:
		newState = t.multiLineCommentState(filepath, line, linecnt, prevToken, currToken, nextToken)
	}

	t.state = newState
}

func (t *Traverser) nonCommentState(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) state.CommentState {
	if currToken == '/' && nextToken == '/' {
		t.buffer += string(currToken)

		return state.SingleLineComment
	} else if currToken == '/' && nextToken == '*' {
		t.buffer += string(currToken)
		t.filepath = filepath
		t.lines = []string{line}
		t.linecnt = linecnt

		return state.MultiLineComment
	} else if currToken == '"' || currToken == '\'' {
		t.stringToken = currToken

		return state.String
	}

	return state.NonComment
}

func (t *Traverser) stringState(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) state.CommentState {
	if prevToken != '\\' && currToken == t.stringToken {
		return state.NonComment
	}

	return state.String
}

func (t *Traverser) singleLineCommentState(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) state.CommentState {
	if currToken == '\n' {
		t.callback(t.buffer, filepath, []string{line}, linecnt)
		t.resetState()

		return state.NonComment
	}

	t.buffer += string(currToken)

	return state.SingleLineComment
}

func (t *Traverser) multiLineCommentState(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) state.CommentState {
	if prevToken == '*' && currToken == '/' {
		t.buffer += string(currToken)

		t.callback(t.buffer, filepath, t.lines, t.linecnt)
		t.resetState()

		return state.NonComment
	}

	if prevToken == '\n' {
		t.lines = append(t.lines, line)
	}

	t.buffer += string(currToken)

	return state.MultiLineComment
}

func (t *Traverser) resetState() {
	t.buffer = ""
	t.filepath = ""
	t.lines = nil
	t.linecnt = 0
}
