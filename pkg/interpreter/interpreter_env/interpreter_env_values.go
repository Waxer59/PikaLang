package interpreter_env

import (
	"pika/pkg/ast"
)

type ValueType string

// Value types
const (
	Null     ValueType = "null"
	Number   ValueType = "number"
	String   ValueType = "string"
	Boolean  ValueType = "boolean"
	Object   ValueType = "object"
	Function ValueType = "function"
	Array    ValueType = "array"
)

type RuntimeValue interface {
	GetType() ValueType
	GetValue() interface{}
}

type NullVal struct {
	Type  ValueType
	Value interface{} // nil
}

func (n NullVal) GetType() ValueType {
	return n.Type
}

func (n NullVal) GetValue() interface{} {
	return n.Value
}

type BooleanVal struct {
	Type  ValueType
	Value bool
}

func (b BooleanVal) GetType() ValueType {
	return b.Type
}

func (b BooleanVal) GetValue() interface{} {
	return b.Value
}

type NumberVal struct {
	Type  ValueType
	Value float64
}

func (n NumberVal) GetType() ValueType {
	return n.Type
}

func (n NumberVal) GetValue() interface{} {
	return n.Value
}

type ObjectVal struct {
	Type       ValueType
	Properties map[string]RuntimeValue
}

func (o ObjectVal) GetType() ValueType {
	return o.Type
}

func (o ObjectVal) GetValue() interface{} {
	return o.Properties
}

type FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue

type FunctionVal struct {
	Type           ValueType
	Name           string
	Params         []string
	DeclarationEnv *Environment
	Body           []ast.Stmt
}

func (f FunctionVal) GetType() ValueType {
	return f.Type
}

func (f FunctionVal) GetValue() interface{} {
	return f.Body
}

type StringVal struct {
	Type  ValueType
	Value string
}

func (s StringVal) GetType() ValueType {
	return s.Type
}

func (s StringVal) GetValue() interface{} {
	return s.Value
}

type NaNVal struct {
	Type  ValueType
	Value string
}

func (n NaNVal) GetType() ValueType {
	return n.Type
}

func (n NaNVal) GetValue() interface{} {
	return n.Value
}

type ArrayVal struct {
	Type     ValueType
	Elements []RuntimeValue
}

func (a ArrayVal) GetType() ValueType {
	return a.Type
}

func (a ArrayVal) GetValue() interface{} {
	return a.Elements
}

func (a ArrayVal) GetElements() []interface{} {
	arr := make([]interface{}, len(a.Elements))
	for idx, element := range a.Elements {
		arr[idx] = element.GetValue()
	}

	return arr
}
