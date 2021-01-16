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

	l.eatWhiteSpace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// initialize tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// return true or false based on whether the character is an indentifier
func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_' || ch == '!' || ch == '?'
}

// read identifier and advance lexer position until non-letter encountered
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.Input[position:l.position]
}

// skip over whitespaces, tab, new line and return
func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
