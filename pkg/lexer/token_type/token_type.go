package token_type

type TokenType int

const (
	// Literal types
	Number TokenType = iota
	Identifier

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
}

var SkippableChars = []string{" ", "\t", "\n", "\r"}

type Token struct {
	Value string
	Type  TokenType
}
