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

	variable, err := env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)

	return variable, err
}

func evalIfStatement(declaration ast.IfStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	conditionRawValue, err := Evaluate(declaration.Condition, env)

	if err != nil {
		return nil, err
	}

	val := conditionRawValue.GetValue()

	// Handle truthy/falsy values
	switch val.(type) {
	case string:
		val = val.(string) != ""
	case int, float32, float64:
		val = val != 0
	default:
		val = false
	}

	if val.(bool) {
		for _, statement := range declaration.Body {
			Evaluate(statement, env)
		}
	}

	return nil, nil
}

func evalFunctionDeclaration(declaration ast.FunctionDeclaration, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {

	fn := interpreter_env.FunctionVal{
		Type:           interpreter_env.Function,
		Name:           declaration.Name,
		Params:         declaration.Params,
		DeclarationEnv: &env,
		Body:           declaration.Body,
	}

	fnName, err := env.DeclareVar(declaration.Name, fn, true)

	return fnName, err
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
