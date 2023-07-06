package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"

	"golang.org/x/exp/slices"
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

func evalSwitchStatement(declaration ast.SwitchStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	for _, caseStatement := range declaration.CaseStmts {

		if slices.ContainsFunc(caseStatement.Test, func(expr ast.Expr) bool {
			eval, err := Evaluate(expr, env)
			return err == nil && eval.GetValue() == declaration.Discriminant || eval.GetValue() == true
		}) {
			for _, statement := range caseStatement.Body {
				_, err := Evaluate(statement, env)
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		}
	}

	if declaration.DefaultStmt.Body == nil {
		return nil, nil
	}

	for _, statement := range declaration.DefaultStmt.Body {
		_, err := Evaluate(statement, env)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func evalIfStatement(declaration ast.IfStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	conditionRawValue, err := Evaluate(declaration.Test, env)

	if err != nil {
		return nil, err
	}

	val := EvaluateTruthyFalsyValues(conditionRawValue.GetValue())

	if err != nil {
		return nil, err
	}

	// Handle first if
	if val {
		for _, statement := range declaration.Body {
			_, err := Evaluate(statement, env)

			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	if declaration.ElseIfStmt == nil && declaration.ElseBody == nil {
		return nil, nil
	}

	// Handle else if
	for _, elseIfStatement := range declaration.ElseIfStmt {
		conditionRawValue, err := Evaluate(elseIfStatement.Test, env)

		if err != nil {
			return nil, err
		}

		val := EvaluateTruthyFalsyValues(conditionRawValue.GetValue())

		if err != nil {
			return nil, err
		}

		if val {
			for _, statement := range elseIfStatement.Body {
				_, err := Evaluate(statement, env)

				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		}
	}

	// Handle else
	for _, statement := range declaration.ElseBody {
		_, err := Evaluate(statement, env)

		if err != nil {
			return nil, err
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
