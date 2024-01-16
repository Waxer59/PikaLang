package compilerErrors

const (
	ErrSyntaxExpectedLeftParen            = "ERROR: Expected '('"
	ErrSyntaxExpectedRightParen           = "ERROR: Expected ')'"
	ErrSyntaxExpectedRightBracket         = "ERROR: Expected ']'"
	ErrSyntaxExpectedLeftBrace            = "ERROR: Expected '{'"
	ErrSyntaxExpectedRightBrace           = "ERROR: Expected '}'"
	ErrSyntaxExpectedColon                = "ERROR: Expected ':'"
	ErrSyntaxExpectedComma                = "ERROR: Expected ','"
	ErrSyntaxExpectedSemicolon            = "ERROR: Expected ';'"
	ErrSyntaxExpectedDoubleQuote          = "ERROR: Expected '\"'"
	ErrSyntaxExpectedIdentifier           = "ERROR: Expected identifier"
	ErrSyntaxInvalidAssignment            = "ERROR: Invalid assignment"
	ErrSyntaxExpectedKey                  = "ERROR: Expected a key"
	ErrSyntaxExpectedAssignation          = "ERROR: Expected '='"
	ErrSyntaxUnterminatedMultilineComment = "ERROR: Unterminated multiline comment"
	ErrConditionCannotBeEmpty             = "ERROR: Condition cannot be empty"
	ErrSyntaxCaseCannotBeEmpty            = "ERROR: Case cannot be empty"
	ErrSyntaxUnaryInvalidUnaryExpr        = "ERROR: Invalid unary expression"
	ErrSyntaxInvalidUpdateExpr            = "ERROR: Invalid update expression"
	ErrSyntaxConditionCantBeEmpty         = "ERROR: Condition cannot be empty"
	ErrParsingError                       = "ERROR: Parsing error"
)
