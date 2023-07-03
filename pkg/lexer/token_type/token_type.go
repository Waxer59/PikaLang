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
	ElseIf

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
	"else if": ElseIf,
}

var SkippableChars = []rune{' ', '\t', '\n', '\r'}

type Token struct {
	Value string
	Type  TokenType
}
