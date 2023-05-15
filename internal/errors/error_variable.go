package errors

import "errors"

type ErrVariable error

var (
	ErrVariableNotFound ErrVariable = errors.New("variable not found")
)
