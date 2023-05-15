package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

func Evaluate(astNode ast.Stmt, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	switch astNode.GetKind() {

	// LITERALS
	case ast_types.Identifier:
		Identifier, err := evalIdentifier(astNode.(ast.Identifier), env)
		return Identifier, err
	case ast_types.NumericLiteral:
		value := astNode.(ast.NumericLiteral).GetValue().(int)
		return interpreter_env.NumberVal{Value: value, Type: interpreter_env.Number}, nil
	case ast_types.ObjectLiteral:
		obj, err := evalObjectExpr(astNode.(ast.ObjectLiteral), env)
		return obj, err
	case ast_types.NullLiteral:
		return interpreter_makers.MK_NULL(), nil
	case ast_types.BooleanLiteral:
		return interpreter_makers.MK_Boolean(astNode.(ast.BooleanLiteral).Value), nil
	case ast_types.StringLiteral:
		return interpreter_makers.MK_String(astNode.(ast.StringLiteral).Value), nil

	// EXPRESSIONS
	case ast_types.BinaryExpr:
		binaryExpr, err := evalBinaryExpr(astNode.(ast.BinaryExpr), env)
		return binaryExpr, err
	case ast_types.CallExpr:
		callExpr, err := evalCallExpr(astNode.(ast.CallExpr), env)
		return callExpr, err
	case ast_types.AssigmentExpr:
		return evalAssignment(astNode.(ast.AssigmentExpr), env)

	// STATEMENTS
	case ast_types.Program:
		program, err := evalProgram(astNode.(ast.Program), env)
		return program, err
	case ast_types.VariableDeclaration:
		variable, err := evalVariableDeclaration(astNode.(ast.VariableDeclaration), env)
		return variable, err
	case ast_types.FunctionDeclaration:
		return evalFunctionDeclaration(astNode.(ast.FunctionDeclaration), env), nil

	default:
		panic("This AST node is not supported")
	}
}
