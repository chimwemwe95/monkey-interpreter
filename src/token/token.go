package token

/*
We define TokenType to be  a string, which allows us to use
many different values as token types. String is the right type
for this book, although int and byte may be more performant
*/
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
Limited number of dfferent tokentypes in the Monkey language
*/

const (
	// special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y...
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
