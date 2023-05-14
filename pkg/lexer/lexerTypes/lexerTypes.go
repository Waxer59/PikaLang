package lexerTypes

type TokenType int

const (
	// Literal types
	Number TokenType = iota
	Identifier

	// Keywords
	Var
	Const
	Func

	// Operators
	BinaryOperator // + - * /
	Equals         // =

	// Grouping
	RightParen // )
	LeftParen  // (
	RightBrace // }
	LeftBrace  // {
	Colon      // :
	SemiColon  // ;
	Comma      // ,

	// End Of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"var":   Var,
	"const": Const,
	"func":  Func,
}

var SkippableChars = []string{" ", "\t", "\n", "\r"}

type Token struct {
	Value string
	Type  TokenType
}
