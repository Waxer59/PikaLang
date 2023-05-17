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

	// Operators
	BinaryOperator // + - * /
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

	// End Of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"var":   Var,
	"const": Const,
	"fn":    Fn,
	"null":  Null,
	"true":  BooleanLiteral,
	"false": BooleanLiteral,
}

var SkippableChars = []rune{' ', '\t', '\n', '\r'}

type Token struct {
	Value string
	Type  TokenType
}
