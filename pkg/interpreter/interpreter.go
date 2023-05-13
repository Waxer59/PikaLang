package interpreter

import (
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/interpreter/interpreterValues"
)

func Evaluate(astNode ast.Stmt) interpreterValues.RuntimeValue {
	switch astNode.GetKind() {
	case astTypes.NumericLiteral:
		value := astNode.(*ast.NumericLiteral).GetValue().(int)
		return interpreterValues.NumberVal{Value: value, Type: "number"}
	case astTypes.NullLiteral:
		return interpreterValues.NullVal{Type: "null", Value: "null"}
	default:
		panic("This AST node is not supported")
	}
}
