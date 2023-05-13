package interpreterEnvironment

import (
	"fmt"
	"pika/pkg/interpreter/interpreterValues"
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

func (e *Environment) DeclareVar(varName string, value interpreterValues.RuntimeValue, constant bool) interpreterValues.RuntimeValue {
	if _, ok := e.variables[varName]; ok {
		panic("Variable already exists")
	}

	e.variables[varName] = value

	if constant {
		e.constants[varName] = value
	}

	return value
}

func (e *Environment) AssignVar(varName string, value interpreterValues.RuntimeValue) interpreterValues.RuntimeValue {
	if _, ok := e.constants[varName]; ok {
		panic("Cannot reassign to a constant: " + varName)
	}

	env := e.Resolve(varName)
	env.variables[varName] = value

	return value
}

// Returns the environment that contains the variable
func (e *Environment) Resolve(varName string) Environment {
	fmt.Println(*e)
	if _, ok := e.variables[varName]; ok {
		return *e
	}

	if e.parent == nil {
		panic("Cannot resolve variable " + varName + " as it does not exist in this scope")
	}

	return e.parent.Resolve(varName)
}

func (e *Environment) LookupVar(varName string) interpreterValues.RuntimeValue {
	env := e.Resolve(varName)
	return env.variables[varName].(interpreterValues.RuntimeValue)
}