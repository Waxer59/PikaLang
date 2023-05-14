package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/interpreter/interpreter_env"
)

func Evaluate(astNode ast.Stmt, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	switch astNode.GetKind() {
	case ast_types.NumericLiteral:
		value := astNode.(ast.NumericLiteral).GetValue().(int)
		return interpreter_env.NumberVal{Value: value, Type: interpreter_env.Number}
	case ast_types.BinaryExpr:
		return evalBinaryExpr(astNode.(ast.BinaryExpr), env)
	case ast_types.Program:
		return evalProgram(astNode.(ast.Program), env)
	case ast_types.Identifier:
		return evalIdentifier(astNode.(ast.Identifier), env)
	case ast_types.VariableDeclaration:
		return evalVariableDeclaration(astNode.(ast.VariableDeclaration), env)
	case ast_types.ObjectLiteral:
		return evalObjectExpr(astNode.(ast.ObjectLiteral), env)
	case ast_types.CallExpr:
		return evalCallExpr(astNode.(ast.CallExpr), env)
	case ast_types.FunctionDeclaration:
		return evalFunctionDeclaration(astNode.(ast.FunctionDeclaration), env)
	case ast_types.AssigmentExpr:
		return evalAssignment(astNode.(ast.AssigmentExpr), env)
	default:
		panic("This AST node is not supported")
	}
}