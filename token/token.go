package token

type TokenType string

type Token struct {
	Type    TokenType //what kind of token is it (maybe its an indentifier/integer)
	Literal string    //Literal Value of the Token
}

// These are just the token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" //End of File

	IDENT = "IDENT" //All user defined identifiers
	INT   = "INT"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok { //Just Checks if the Given Identifier is a Keyword
		return tok //returns the tokentype if it is infact a keyword
	}
	return IDENT //otherwise its a userdeined identi
}
