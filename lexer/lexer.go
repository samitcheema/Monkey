package lexer

import "github.com/samitcheema/Monkey/token"

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

// NextToken ...
func (l *Lexer) NextToken() token.Token {
	var tok token.Token //variable which will hold the token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RPAREN, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
