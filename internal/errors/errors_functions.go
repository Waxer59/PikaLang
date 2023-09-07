package compilerErrors

const (
	ErrFuncNotFound                 = "ERROR: Function not found: "
	ErrFuncExpectedIdentifer        = "ERROR: Expected identifier for function name"
	ErrReturn                       = "ERROR: Return statement outside of function"
	ErrNotEnoughArguments           = "ERROR: Not enough arguments for function"
	ErrTooManyArguments             = "ERROR: Too many arguments for function"
	ErrUndefinedFunction            = "ERROR: Undefined function: "
	ErrComputedPropertyMustBeString = "ERROR: Computed property must be a string"
)
