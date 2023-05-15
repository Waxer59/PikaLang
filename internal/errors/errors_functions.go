package compilerErrors

type ErrFunc string

var (
	ErrFuncNotFound          ErrFunc = "ERROR: Function not found: "
	ErrFuncExpectedIdentifer ErrFunc = "ERROR: Expected identifier for function name"
)
