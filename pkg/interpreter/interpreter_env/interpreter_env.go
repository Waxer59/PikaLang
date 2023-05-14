package interpreter_env

import (
	"fmt"
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

func (e *Environment) DeclareVar(varName string, value RuntimeValue, constant bool) RuntimeValue {
	fmt.Println("Declaring variable " + varName)

	if _, ok := e.variables[varName]; ok {
		panic("Variable already exists")
	}

	e.variables[varName] = value

	if constant {
		e.constants[varName] = value
	}

	return value
}

func (e *Environment) AssignVar(varName string, value RuntimeValue) RuntimeValue {
	if _, ok := e.constants[varName]; ok {
		panic("Cannot reassign to a constant: " + varName)
	}

	env := e.Resolve(varName)
	env.variables[varName] = value

	return value
}

// Returns the environment that contains the variable
func (e *Environment) Resolve(varName string) Environment {
	if _, ok := e.variables[varName]; ok {
		return *e
	}

	if e.parent == nil {
		panic("Cannot resolve variable " + varName + " as it does not exist in this scope")
	}

	return e.parent.Resolve(varName)
}

func (e *Environment) LookupVar(varName string) RuntimeValue {
	env := e.Resolve(varName)
	return env.variables[varName].(RuntimeValue)
}
