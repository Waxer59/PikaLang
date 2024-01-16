package compilerErrors

const (
	ErrVariableDoesNotExist                              = "ERROR: Variable does not exist in this scope: "
	ErrVariableAlreadyExists                             = "ERROR: Variable already exists: "
	ErrVariableIsConstant                                = "ERROR: Constant cant be re-assigned: "
	ErrVariableExpectedIdentifierNameFollowingConstOrVar = "ERROR: Expected identifier name following 'const' or 'var'"
)
