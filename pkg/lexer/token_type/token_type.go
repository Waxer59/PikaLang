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
	NaN
	Return
	While
	Continue
	Break
	For

	// Operators
	BinaryOperator // + - * / ** %

	// Assigment operators
	Equals       // =
	PlusEquals   // +=
	MinusEquals  // -=
	TimesEquals  // *=
	DivideEquals // /=
	PowerEquals  // **=
	ModuleEquals // %=
	Arrow        // =>

	// Update operators
	Increment // ++
	Decrement // --

	// Grouping
	LeftParen    // (
	RightParen   // )
	RightBrace   // }
	LeftBrace    // {
	RightBracket // ]
	LeftBracket  // [
	Colon        // :
	Semicolon    // ;
	Comma        // ,
	Dot          // .
	DoubleQuote  // "
	SingleQuote  // '
	QuestionMark // ?

	// Comparison operators
	EqualEqual   // ==
	Greater      // >
	GreaterEqual // >=
	Less         // <
	LessEqual    // <=
	Bang         // !
	NotEqual     // !=
	Or           // ||
	And          // &&

	// End Of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"var":      Var,
	"const":    Const,
	"fn":       Fn,
	"null":     Null,
	"true":     BooleanLiteral,
	"false":    BooleanLiteral,
	"if":       If,
	"else":     Else,
	"switch":   Switch,
	"case":     Case,
	"default":  Default,
	"NaN":      NaN,
	"return":   Return,
	"while":    While,
	"continue": Continue,
	"break":    Break,
	"for":      For,
}

var SkippableChars = []rune{' ', '\t', '\n', '\r'}

// IDENTIFIERS
var AllowedIdentifierChars = []rune{'_', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var AllowedIdentifierCharsWithFirst = []rune{'_'}

// Assigment operators
var AssigmentOperators = []TokenType{Equals, PlusEquals, MinusEquals, TimesEquals, DivideEquals, PowerEquals, ModuleEquals}

type Token struct {
	Value string
	Type  TokenType
}
