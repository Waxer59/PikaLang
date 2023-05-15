package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
)

func evalFunctionDeclaration(declaration ast.FunctionDeclaration, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {

	fn := interpreter_env.FunctionVal{
		Type:           interpreter_env.Function,
		Name:           declaration.Name,
		Params:         declaration.Params,
		DeclarationEnv: &env,
		Body:           declaration.Body,
	}

	variables, err := env.DeclareVar(declaration.Name, fn, true)

	return variables, err
}
