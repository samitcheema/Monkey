package token

type TokenType string // Allow us to use different values as TokenType

type Token struct {
	Type    TokenType
	Literal string
}

// map to differentiate from
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Function checks whether identifier is user defined or a language keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT" //Integers

	// Special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Operators
	ASSIGN = "="
	ADD    = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = "("
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
