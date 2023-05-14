package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

func evalVariableDeclaration(variableDeclaration ast.VariableDeclaration, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	var value interpreter_env.RuntimeValue = interpreter_makers.MK_NULL()

	if variableDeclaration.Value != nil {
		value = Evaluate(variableDeclaration.Value, env)
	}
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)
}

func evalProgram(program ast.Program, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	var lastEvaluated interpreter_env.RuntimeValue = interpreter_makers.MK_NULL()

	for _, statement := range program.Body {
		lastEvaluated = Evaluate(statement, env)
	}

	return lastEvaluated
}
