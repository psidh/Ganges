package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Following are the different token types

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords in Sankrit (for the most part)
	KRIYA   = "KRIYA"
	RAMA    = "RAMA"
	SATYA   = "SATYA"
	ASATYA  = "ASATYA"
	YADI    = "YADI"
	ANYATHA = "ANYATHA"
	DAAN    = "DAAN"
	VAKYA   = "VAKYA"
	CHAKRA  = "CHAKRA"
)

var keywords = map[string]TokenType{
	"kriya":   KRIYA,
	"rama":    RAMA,
	"yadi":    YADI,
	"daan":    DAAN,
	"anyatha": ANYATHA,
	"satya":   SATYA,
	"asatya":  ASATYA,
	"chakra":  CHAKRA,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
