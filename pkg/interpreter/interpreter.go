package interpreter

import (
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/interpreter/interpreterValues"
)

func evaluateBinaryExpr(binop ast.BinaryExpr) interpreterValues.RuntimeValue {
	lhs := Evaluate(binop.Left)
	rhs := Evaluate(binop.Right)

	if lhs.GetType() == interpreterValues.Number && rhs.GetType() == interpreterValues.Number {
		return evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
	}

	return interpreterValues.NullVal{Type: "null", Value: "null"}
}

func evaluateNumericBinaryExpr(operator string, lhs interpreterValues.RuntimeValue, rhs interpreterValues.RuntimeValue) interpreterValues.RuntimeValue {
	result := 0

	valLhs, okLhs := lhs.(interpreterValues.NumberVal)
	valRhs, okRhs := rhs.(interpreterValues.NumberVal)

	if !okLhs || !okRhs {
		panic("Left and right hand side of binary expression must be of type number")
	}

	switch operator {
	case "+":
		result = valLhs.Value + valRhs.Value
	case "-":
		result = valLhs.Value - valRhs.Value
	case "*":
		result = valLhs.Value * valRhs.Value
	case "/":
		//TODO: Division by zero checks
		result = valLhs.Value / valRhs.Value
	case "%":
		result = valLhs.Value % valRhs.Value
	}

	return interpreterValues.NumberVal{Value: result, Type: interpreterValues.Number}
}

func evaluateProgram(program ast.Program) interpreterValues.RuntimeValue {
	var lastEvaluated interpreterValues.RuntimeValue = interpreterValues.NullVal{Type: "null", Value: "null"}

	for _, statement := range program.Body {
		lastEvaluated = Evaluate(statement)
	}

	return lastEvaluated
}

func Evaluate(astNode ast.Stmt) interpreterValues.RuntimeValue {
	switch astNode.GetKind() {
	case astTypes.NumericLiteral:
		value := astNode.(ast.NumericLiteral).GetValue().(int)
		return interpreterValues.NumberVal{Value: value, Type: "number"}
	case astTypes.BinaryExpr:
		return evaluateBinaryExpr(astNode.(ast.BinaryExpr))
	case astTypes.Program:
		return evaluateProgram(astNode.(ast.Program))
	case astTypes.NullLiteral:
		return interpreterValues.NullVal{Type: "null", Value: "null"}
	default:
		panic("This AST node is not supported")
	}
}
