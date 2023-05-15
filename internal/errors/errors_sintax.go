package compilerErrors

type ErrSintax string

var (
	ErrSintaxExpectedLeftParen    ErrSintax = "ERROR: Expected '('"
	ErrSintaxExpectedRightParen   ErrSintax = "ERROR: Expected ')'"
	ErrSintaxExpectedLeftBracket  ErrSintax = "ERROR: Expected '['"
	ErrSintaxExpectedRightBracket ErrSintax = "ERROR: Expected ']'"
	ErrSintaxExpectedLeftBrace    ErrSintax = "ERROR: Expected '{'"
	ErrSintaxExpectedRightBrace   ErrSintax = "ERROR: Expected '}'"
	ErrSintaxExpectedColon        ErrSintax = "ERROR: Expected ':'"
	ErrSintaxExpectedComma        ErrSintax = "ERROR: Expected ','"
	ErrSintaxExpectedSemicolon    ErrSintax = "ERROR: Expected ';'"
	ErrSintaxExpectedDot          ErrSintax = "ERROR: Expected '.'"
	ErrSintaxExpectedOperator     ErrSintax = "ERROR: Expected operator"
	ErrSintaxExpectedDoubleQoute  ErrSintax = "ERROR: Expected '\"'"
	ErrSintaxExpectedSingleQuote  ErrSintax = "ERROR: Expected \"'\""
	ErrSintaxExpectedIdentifier   ErrSintax = "ERROR: Expected identifier"
	ErrSintaxInvalidAssignment    ErrSintax = "ERROR: Invalid assignment"
	ErrSintaxExpectedKey          ErrSintax = "ERROR: Expected a key"
	ErrSintaxExpectedAsignation   ErrSintax = "ERROR: Expected '='"
)
