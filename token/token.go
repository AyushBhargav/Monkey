package token

// Type custom token type synonymous with string
type Type string

// Token encapsulates token type with its literal representation
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL token, about which we don't know about.
	ILLEGAL = "ILLEGAL"
	// EOF denotes end of file
	EOF = "EOF"
	// IDENT denotes any identifier like x, y, foobar etc.
	IDENT = "IDENTIFIER"
	// INT denotes integers like 2, 100, -100
	INT = "INT"
	// ASSIGN denotes =
	ASSIGN = "="
	// PLUS denotes +
	PLUS = "+"
	// MINUS denotes -
	MINUS = "-"
	// BANG denotes ! used for negation
	BANG = "!"
	// ASTERISK denotes multiplication etc.
	ASTERISK = "*"
	// SLASH denotes division.
	SLASH = "/"
	// LT denotes less than operation.
	LT = "<"
	// RT denotes greater than operation.
	RT = ">"
	// COMMA denotes ,
	COMMA = ","
	// SEMICOLON denotes statement separator like ,
	SEMICOLON = ";"
	// LPAREN denotes left parenthesis
	LPAREN = "("
	// RPAREN denotes right parenthesis
	RPAREN = ")"
	// LBRACE denotes left brace
	LBRACE = "{"
	// RBRACE denotes right brace
	RBRACE = "}"
	// EQUAL denotes ==
	EQUAL = "=="
	// UNEQUAL denotes !=
	UNEQUAL = "!="
	// FUNCTION denotes 'fn' keyword
	FUNCTION = "FUNCTION"
	// LET denotes 'let' keyword
	LET = "LET"
	// IF denotes if conditional keyword
	IF = "IF"
	// ELSE denotes else conditional keyword
	ELSE = "ELSE"
	// TRUE denotes boolean true
	TRUE = "TRUE"
	// FALSE denotes boolean false
	FALSE = "FALSE"
	// RETURN denotes return keyword
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent returns if identifier is a reserver word, else return argument as it is.
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
