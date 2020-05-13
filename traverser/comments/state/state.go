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
