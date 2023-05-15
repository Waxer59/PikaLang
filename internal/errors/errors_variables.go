package compilerErrors

type ErrVariable string

var (
	ErrVariableDoesNotExist                              ErrVariable = "ERROR: Variable does not exist in this scope: "
	ErrVariableAlreadyExists                             ErrVariable = "ERROR: Variable already exists: "
	ErrVariableIsConstant                                ErrVariable = "ERROR: Constant cant be re-assigned: "
	ErrVariableExpectedIdentifierNameFollowingConstOrVar ErrVariable = "ERROR: Expected identifier name following 'const' or 'var'"
)
