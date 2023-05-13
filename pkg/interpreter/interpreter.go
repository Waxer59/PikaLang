package interpreter

import (
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/interpreter/interpreterEnvironment"
	"pika/pkg/interpreter/interpreterMakers"
	"pika/pkg/interpreter/interpreterValues"
)

func evaluateBinaryExpr(binop ast.BinaryExpr, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	lhs := Evaluate(binop.Left, env)
	rhs := Evaluate(binop.Right, env)

	if lhs.GetType() == interpreterValues.Number && rhs.GetType() == interpreterValues.Number {
		return evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
	}

	return interpreterMakers.MK_NULL()
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

func evaluateProgram(program ast.Program, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	var lastEvaluated interpreterValues.RuntimeValue = interpreterMakers.MK_NULL()

	for _, statement := range program.Body {
		lastEvaluated = Evaluate(statement, env)
	}

	return lastEvaluated
}

func evalIdentifier(ident ast.Identifier, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	val := env.LookupVar(ident.Symbol)
	return val
}

func evalVariableDeclaration(variableDeclaration ast.VariableDeclaration, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	var value interpreterValues.RuntimeValue = interpreterMakers.MK_NULL()

	if variableDeclaration.Value != nil {
		value = Evaluate(variableDeclaration.Value, env)
	}
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)
}

func evalAssignment(assignment ast.AssigmentExpr, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	if assignment.Assigne.GetKind() != astTypes.Identifier {
		panic("Invalid assignment target")
	}

	varName := assignment.Assigne.(ast.Identifier).Symbol
	return env.AssignVar(varName, Evaluate(assignment.Value, env))
}

func Evaluate(astNode ast.Stmt, env interpreterEnvironment.Environment) interpreterValues.RuntimeValue {
	switch astNode.GetKind() {
	case astTypes.NumericLiteral:
		value := astNode.(ast.NumericLiteral).GetValue().(int)
		return interpreterValues.NumberVal{Value: value, Type: interpreterValues.Number}
	case astTypes.BinaryExpr:
		return evaluateBinaryExpr(astNode.(ast.BinaryExpr), env)
	case astTypes.Program:
		return evaluateProgram(astNode.(ast.Program), env)
	case astTypes.Identifier:
		return evalIdentifier(astNode.(ast.Identifier), env)
	case astTypes.VariableDeclaration:
		return evalVariableDeclaration(astNode.(ast.VariableDeclaration), env)
	case astTypes.AssigmentExpr:
		return evalAssignment(astNode.(ast.AssigmentExpr), env)
	default:
		panic("This AST node is not supported")
	}
}
