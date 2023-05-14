package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/interpreter/interpreter_env"
)

func evalFunctionDeclaration(declaration ast.FunctionDeclaration, env interpreter_env.Environment) interpreter_env.RuntimeValue {

	fn := interpreter_env.FunctionVal{
		Type:           interpreter_env.Function,
		Name:           declaration.Name,
		Params:         declaration.Params,
		DeclarationEnv: &env,
		Body:           declaration.Body,
	}

	return env.DeclareVar(declaration.Name, fn, true)
}
