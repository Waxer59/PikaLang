package interpreter_env

import (
	"errors"
	compilerErrors "pika/internal/errors"
)

type Environment struct {
	parent    *Environment
	variables map[string]interface{}
	constants map[string]interface{}
}

func New(parentENV *Environment) Environment {
	return Environment{
		parent:    parentENV,
		variables: make(map[string]interface{}),
		constants: make(map[string]interface{}),
	}
}

func (e *Environment) DeclareVar(varName string, value RuntimeValue, constant bool) (RuntimeValue, error) {

	if _, ok := e.variables[varName]; ok {
		return nil, errors.New(compilerErrors.ErrVariableAlreadyExists + varName)
	}

	e.variables[varName] = value

	if constant {
		e.constants[varName] = value
	}

	return value, nil
}

func (e *Environment) AssignVar(varName string, value RuntimeValue) (RuntimeValue, error) {
	if _, ok := e.constants[varName]; ok {
		return nil, errors.New(compilerErrors.ErrVariableIsConstant + varName)
	}

	env, err := e.Resolve(varName)
	env.variables[varName] = value

	return value, err
}

// Returns the environment that contains the variable
func (e *Environment) Resolve(varName string) (Environment, error) {
	if _, ok := e.variables[varName]; ok {
		return *e, nil
	}

	if e.parent == nil {
		return *e, errors.New(compilerErrors.ErrVariableDoesNotExist + varName)
	}

	return e.parent.Resolve(varName)
}

func (e *Environment) LookupVar(varName string) (RuntimeValue, error) {
	env, err := e.Resolve(varName)
	if err != nil {
		return nil, err
	}
	return env.variables[varName].(RuntimeValue), nil
}
