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
	GetValue() any
}

type NullVal struct {
	Type  ValueType
	Value any // nil
}

func (n NullVal) GetType() ValueType {
	return n.Type
}

func (n NullVal) GetValue() any {
	return n.Value
}

type BooleanVal struct {
	Type  ValueType
	Value bool
}

func (b BooleanVal) GetType() ValueType {
	return b.Type
}

func (b BooleanVal) GetValue() any {
	return b.Value
}

type NumberVal struct {
	Type  ValueType
	Value float64
}

func (n NumberVal) GetType() ValueType {
	return n.Type
}

func (n NumberVal) GetValue() any {
	return n.Value
}

type ObjectVal struct {
	Type       ValueType
	Properties map[string]RuntimeValue
}

func (o ObjectVal) GetType() ValueType {
	return o.Type
}

func (o ObjectVal) GetValue() any {
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

func (f FunctionVal) GetValue() any {
	return f.Body
}

type StringVal struct {
	Type  ValueType
	Value string
}

func (s StringVal) GetType() ValueType {
	return s.Type
}

func (s StringVal) GetValue() any {
	return s.Value
}

type NaNVal struct {
	Type  ValueType
	Value string
}

func (n NaNVal) GetType() ValueType {
	return n.Type
}

func (n NaNVal) GetValue() any {
	return n.Value
}

type ArrayVal struct {
	Type     ValueType
	Elements []RuntimeValue
}

func (a ArrayVal) GetType() ValueType {
	return a.Type
}

func (a ArrayVal) GetValue() any {
	return a.Elements
}

func (a ArrayVal) GetElements() []any {
	arr := make([]any, len(a.Elements))
	for idx, element := range a.Elements {
		arr[idx] = element.GetValue()
	}

	return arr
}
