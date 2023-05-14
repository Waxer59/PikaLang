package interpreter

import "pika/pkg/ast"

type ValueType string

const (
	Null     ValueType = "null"
	Number   ValueType = "number"
	Boolean  ValueType = "boolean"
	Object   ValueType = "object"
	NativeFn ValueType = "nativeFn"
	Function ValueType = "function"
)

type RuntimeValue interface {
	GetType() ValueType
}

type NullVal struct {
	Type  ValueType
	Value interface{} // <-- nil
}

type BooleanVal struct {
	Type  ValueType
	Value bool
}

type NumberVal struct {
	Type  ValueType
	Value int
}

type ObjectVal struct {
	Type       ValueType
	Properties map[string]RuntimeValue
}

type FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue

type NativeFnVal struct {
	Type ValueType
	Call FunctionCall
}

type FunctionVal struct {
	Type           ValueType
	Name           string
	Params         []string
	DeclarationEnv *Environment
	Body           []ast.Stmt
}

func (o ObjectVal) GetType() ValueType {
	return o.Type
}

func (f FunctionVal) GetType() ValueType {
	return f.Type
}

func (n NativeFnVal) GetType() ValueType {
	return n.Type
}

func (n NullVal) GetType() ValueType {
	return n.Type
}

func (b BooleanVal) GetType() ValueType {
	return b.Type
}

func (n NumberVal) GetType() ValueType {
	return n.Type
}
