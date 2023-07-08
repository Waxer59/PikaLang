package compilerErrors

const (
	ErrSyntaxExpectedLeftParen            = "ERROR: Expected '('"
	ErrSyntaxExpectedRightParen           = "ERROR: Expected ')'"
	ErrSyntaxExpectedLeftBracket          = "ERROR: Expected '['"
	ErrSyntaxExpectedRightBracket         = "ERROR: Expected ']'"
	ErrSyntaxExpectedLeftBrace            = "ERROR: Expected '{'"
	ErrSyntaxExpectedRightBrace           = "ERROR: Expected '}'"
	ErrSyntaxExpectedColon                = "ERROR: Expected ':'"
	ErrSyntaxExpectedComma                = "ERROR: Expected ','"
	ErrSyntaxExpectedSemicolon            = "ERROR: Expected ';'"
	ErrSyntaxExpectedDot                  = "ERROR: Expected '.'"
	ErrSyntaxExpectedOperator             = "ERROR: Expected operator"
	ErrSyntaxExpectedDoubleQoute          = "ERROR: Expected '\"'"
	ErrSyntaxExpectedSingleQuote          = "ERROR: Expected \"'\""
	ErrSyntaxExpectedIdentifier           = "ERROR: Expected identifier"
	ErrSyntaxInvalidAssignment            = "ERROR: Invalid assignment"
	ErrSyntaxExpectedKey                  = "ERROR: Expected a key"
	ErrSyntaxExpectedAsignation           = "ERROR: Expected '='"
	ErrSyntaxUnterminatedMultilineComment = "ERROR: Unterminated multiline comment"
	ErrSyntaxStatementNotFound            = "ERROR: Statement not supported"
	ErrSyntaxExpectedExpr                 = "ERROR: Expected an expression"
	ErrConditionCannotBeEmpty             = "ERROR: Condition cannot be empty"
	ErrSyntaxCaseCannotBeEmpty            = "ERROR: Case cannot be empty"
	ErrSyntaxExpectedQuestionMark         = "ERROR: Expected '?'"
	ErrSyntaxUnaryInvalidUnaryExpr        = "ERROR: Invalid unary expression"
)
