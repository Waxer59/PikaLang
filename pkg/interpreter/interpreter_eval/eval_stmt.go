package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

func evalVariableDeclaration(variableDeclaration ast.VariableDeclaration, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	var value interpreter_env.RuntimeValue = interpreter_makers.MK_NULL()

	if variableDeclaration.Value != nil {
		eval, err := Evaluate(variableDeclaration.Value, env)
		if err != nil {
			return nil, err
		}
		value = eval
	}
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant), nil
}

func evalProgram(program ast.Program, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	var lastEvaluated interpreter_env.RuntimeValue = interpreter_makers.MK_NULL()

	for _, statement := range program.Body {
		eval, err := Evaluate(statement, env)
		if err != nil {
			return nil, err
		}
		lastEvaluated = eval
	}

	return lastEvaluated, nil
}
