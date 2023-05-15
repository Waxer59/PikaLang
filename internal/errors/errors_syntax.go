package compilerErrors

type ErrSyntax string

var (
	ErrSyntaxExpectedLeftParen            ErrSyntax = "ERROR: Expected '('"
	ErrSyntaxExpectedRightParen           ErrSyntax = "ERROR: Expected ')'"
	ErrSyntaxExpectedLeftBracket          ErrSyntax = "ERROR: Expected '['"
	ErrSyntaxExpectedRightBracket         ErrSyntax = "ERROR: Expected ']'"
	ErrSyntaxExpectedLeftBrace            ErrSyntax = "ERROR: Expected '{'"
	ErrSyntaxExpectedRightBrace           ErrSyntax = "ERROR: Expected '}'"
	ErrSyntaxExpectedColon                ErrSyntax = "ERROR: Expected ':'"
	ErrSyntaxExpectedComma                ErrSyntax = "ERROR: Expected ','"
	ErrSyntaxExpectedSemicolon            ErrSyntax = "ERROR: Expected ';'"
	ErrSyntaxExpectedDot                  ErrSyntax = "ERROR: Expected '.'"
	ErrSyntaxExpectedOperator             ErrSyntax = "ERROR: Expected operator"
	ErrSyntaxExpectedDoubleQoute          ErrSyntax = "ERROR: Expected '\"'"
	ErrSyntaxExpectedSingleQuote          ErrSyntax = "ERROR: Expected \"'\""
	ErrSyntaxExpectedIdentifier           ErrSyntax = "ERROR: Expected identifier"
	ErrSyntaxInvalidAssignment            ErrSyntax = "ERROR: Invalid assignment"
	ErrSyntaxExpectedKey                  ErrSyntax = "ERROR: Expected a key"
	ErrSyntaxExpectedAsignation           ErrSyntax = "ERROR: Expected '='"
	ErrSyntaxUnterminatedMultilineComment ErrSyntax = "ERROR: Unterminated multiline comment"
)
