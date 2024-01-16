package interpreter_eval

import (
	"errors"

	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/ast/ast_types"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

func Evaluate(astNode ast.Stmt, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	switch astNode.GetKind() {

	// LITERALS
	case ast_types.Identifier:
		return evalIdentifier(astNode.(ast.Identifier), env)
	case ast_types.NumericLiteral:
		return interpreter_makers.MkNumber(astNode.(ast.NumericLiteral).Value), nil
	case ast_types.ObjectLiteral:
		return evalObjectExpr(astNode.(ast.ObjectLiteral), env)
	case ast_types.NullLiteral:
		return interpreter_makers.MkNull(), nil
	case ast_types.BooleanLiteral:
		return interpreter_makers.MkBoolean(astNode.(ast.BooleanLiteral).Value), nil
	case ast_types.StringLiteral:
		return interpreter_makers.MkString(astNode.(ast.StringLiteral).Value), nil
	case ast_types.NaNLiteral:
		return interpreter_makers.MkNan(), nil
	case ast_types.ArrayLiteral:
		return evalArrayExpr(astNode.(ast.ArrayLiteral), env)

	// EXPRESSIONS
	case ast_types.BinaryExpr:
		return evalBinaryExpr(astNode.(ast.BinaryExpr), env)
	case ast_types.CallExpr:
		return evalCallExpr(astNode.(ast.CallExpr), env)
	case ast_types.AssigmentExpr:
		return evalAssignment(astNode.(ast.AssigmentExpr), env)
	case ast_types.ConditionalExpr:
		return evalConditionalExpr(astNode.(ast.ConditionalExpr), env)
	case ast_types.LogicalExpr:
		return evalLogicalExpr(astNode.(ast.LogicalExpr), env)
	case ast_types.UnaryExpr:
		return evalUnaryExpr(astNode.(ast.UnaryExpr), env)
	case ast_types.MemberExpr:
		return evalMemberExpr(astNode.(ast.MemberExpr), env)
	case ast_types.UpdateExpr:
		return evalUpdateExpr(astNode.(ast.UpdateExpr), env)
	case ast_types.ArrowFunctionExpr:
		return evalArrowFunctionExpr(astNode.(ast.ArrowFunctionExpr))

	// STATEMENTS
	case ast_types.Program:
		return evalProgram(astNode.(ast.Program), env)
	case ast_types.VariableDeclaration:
		return evalVariableDeclaration(astNode.(ast.VariableDeclaration), env)
	case ast_types.FunctionDeclaration:
		return evalFunctionDeclaration(astNode.(ast.FunctionDeclaration), env)
	case ast_types.IfStatement:
		return evalIfStatement(astNode.(ast.IfStatement), env)
	case ast_types.SwitchStatement:
		return evalSwitchStatement(astNode.(ast.SwitchStatement), env)
	case ast_types.ReturnStatement:
		return evalReturnStatement(astNode.(ast.ReturnStatement), env)
	case ast_types.WhileStatement:
		return evalWhileStatement(astNode.(ast.WhileStatement), env)
	case ast_types.BreakStatement:
		return evalBreakStatement()
	case ast_types.ContinueStatement:
		return evalContinueStatement()
	case ast_types.ForStatement:
		return evalForStatement(astNode.(ast.ForStatement), env)

	default:
		return nil, errors.New("ERROR: Unknown node type")
	}
}
