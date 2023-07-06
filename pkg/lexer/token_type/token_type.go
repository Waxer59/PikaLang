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
	QuestionMark // ?

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

// IDENTIFIERS
var AllowedIdentifierChars = []rune{'_', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var AllowedIdentifierCharsWithFirst = []rune{'_'}

type Token struct {
	Value string
	Type  TokenType
}
