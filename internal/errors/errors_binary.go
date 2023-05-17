package compilerErrors

type ErrBinary string

var (
	ErrBinaryInvalidBinaryExpr ErrBinary = "ERROR: Invalid binary operation"
	ErrBinaryDivisionByZero    ErrBinary = "ERROR: Division by zero"
)
