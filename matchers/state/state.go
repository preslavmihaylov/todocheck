package state

// CommentState is an enum representing the current state the traverser is in
type CommentState int

const (
	// NonComment is the default state a traverser is at
	NonComment CommentState = iota

	// String is the state while the traverser is reading a string
	String

	// SingleLineComment is the state while the traverser is reading a single-line comment
	SingleLineComment

	// MultiLineComment is the state while the traverser is reading a multi-line comment
	MultiLineComment
)

// Func represents a state-transitioning function used in the state-machine design pattern
// Its parameters are tailored to traversing a file's stream of tokens
type Func func(filepath, line string, linecnt int, prevToken, currToken, nextToken rune) (CommentState, error)

// CommentCallback is a function which acts on an encountered comment
type CommentCallback func(comment, filename string, lines []string, linecnt int) error
