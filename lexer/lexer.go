package lexer

// Lexer group
type Lexer struct {
	Input        string
	position     int  // pointing to the current position within input being examined
	readPosition int  // pointing to the next position after the current one(position)
	ch           byte // the character which is being examined
}

// New function create and point to lexer
func New(input string) *Lexer {
	l := &Lexer{Input: input}
	l.readChar()
	return l
}

// readChar reads character within input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.Input) {
		l.ch = 0
	} else {
		l.ch = l.Input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
