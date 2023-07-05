package token_type

type TokenType int

const (
	// Literal types
	Number TokenType = iota
	Identifier
	Null
	BooleanLiteral
	StringLiteral

	// Keywords
	Var
	Const
	Fn
	If
	Else
	Switch
	Case
	Default

	// Operators
	BinaryOperator // + - * / ** %
	Equals         // =

	// Grouping
	LeftParen    // (
	RightParen   // )
	RightBrace   // }
	LeftBrace    // {
	RightBracket // ]
	LeftBracket  // [
	Colon        // :
	SemiColon    // ;
	Comma        // ,
	Dot          // .
	DoubleQoute  // "
	SingleQoute  // '

	// Comparison operators
	EqualEqual   // ==
	Greater      // >
	GreaterEqual // >=
	Less         // <
	LessEqual    // <=
	Not          // !
	NotEqual     // !=

	// End Of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"var":     Var,
	"const":   Const,
	"fn":      Fn,
	"null":    Null,
	"true":    BooleanLiteral,
	"false":   BooleanLiteral,
	"if":      If,
	"else":    Else,
	"switch":  Switch,
	"case":    Case,
	"default": Default,
}

var SkippableChars = []rune{' ', '\t', '\n', '\r'}
var AllowedIdentifierChars = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '_', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type Token struct {
	Value string
	Type  TokenType
}
