package interpreter_eval

import (
	"errors"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"

	"golang.org/x/exp/slices"
)

func evalVariableDeclaration(variableDeclaration ast.VariableDeclaration, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	var value interpreter_env.RuntimeValue = interpreter_makers.MK_Null()

	if variableDeclaration.Value != nil {
		eval, err := Evaluate(variableDeclaration.Value, env)
		if err != nil {
			return eval, err
		}
		value = eval
	}

	variable, err := env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)

	return variable, err
}

func evalReturnStatement(declaration ast.ReturnStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {

	var returnValue interpreter_env.RuntimeValue = interpreter_makers.MK_Null()

	if declaration.Argument != nil {
		eval, err := Evaluate(declaration.Argument, env)
		if err != nil {
			return eval, err
		}
		returnValue = eval
	}

	// Throw a error to stop the execution
	return returnValue, errors.New(compilerErrors.ErrReturn)
}

func evalBreakStatement(declaration ast.BreakStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	return nil, errors.New(compilerErrors.ErrLoopsBreakNotInLoop)
}

func evalContinueStatement(declaration ast.ContinueStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	return nil, errors.New(compilerErrors.ErrLoopsContinueNotInLoop)
}

func evalWhileStatement(declaration ast.WhileStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	testEval, err := Evaluate(declaration.Test, env)
	if err != nil {
		return nil, err
	}

	testVal := EvaluateTruthyFalsyValues(testEval.GetValue())

	for testVal {
		eval, err := EvaluateBodyStmt(declaration.Body, env)
		if err != nil && err.Error() == compilerErrors.ErrLoopsBreakNotInLoop {
			break
		} else if err != nil && err.Error() == compilerErrors.ErrLoopsContinueNotInLoop {
			continue
		} else if err != nil {
			return eval, err
		}
		testEval, err = Evaluate(declaration.Test, env)

		if err != nil {
			return eval, err
		}

		testVal = EvaluateTruthyFalsyValues(testEval.GetValue())
	}

	return interpreter_makers.MK_Null(), nil
}

func evalSwitchStatement(declaration ast.SwitchStatement, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	for _, caseStatement := range declaration.CaseStmts {

		if slices.ContainsFunc(caseStatement.Test, func(expr ast.Expr) bool {
			eval, err := Evaluate(expr, env)
			if err != nil {
				return false
			}
			evalDiscriminant, err := Evaluate(declaration.Discriminant, env)
			if err != nil {
				return false
			}
			return err == nil && eval.GetValue() == evalDiscriminant.GetValue() || eval.GetValue() == true
		}) {
			eval, err := EvaluateBodyStmt(caseStatement.Body, env)
			if err != nil {
				return eval, err
			}
			return nil, nil
		}
	}

	if declaration.DefaultStmt.Body == nil {
		return nil, nil
	}

	eval, err := EvaluateBodyStmt(declaration.DefaultStmt.Body, env)
	if err != nil {
		return eval, err
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
		eval, err := EvaluateBodyStmt(declaration.Body, env)
		if err != nil {
			return eval, err
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

		if val {
			eval, err := EvaluateBodyStmt(elseIfStatement.Body, env)
			if err != nil {
				return eval, err
			}
			return nil, nil
		}
	}

	// Handle else
	eval, err := EvaluateBodyStmt(declaration.ElseBody, env)
	if err != nil {
		return eval, err
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
	return EvaluateBodyStmt(program.Body, env)
}
