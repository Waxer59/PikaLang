package compilerErrors

type ErrBinary string

var (
	ErrSyntaxInvalidBinaryExpr ErrBinary = "ERROR: Invalid binary operation"
)
