package lexer

// Lexer group
type Lexer struct {
	input        string
	position     int  // pointing to the current position within input being examined
	readPosition int  // pointing to the next position after the current one(position)
	ch           byte // the character which is being examined
}

// New function...
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
