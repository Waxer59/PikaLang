package interpreter_eval

import (
	"errors"
	compilerErrors "github.com/Waxer59/PikaLang/internal/errors"
	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_eval/internal/nativeFns"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

func GetFunctionName(caller ast.CallExpr, env interpreter_env.Environment) (string, error) {
	switch caller.Caller.(type) {
	case ast.Identifier:
		return caller.Caller.(ast.Identifier).Symbol, nil
	case ast.MemberExpr:
		isComputed := caller.Caller.(ast.MemberExpr).Computed

		if isComputed {
			eval, err := Evaluate(caller.Caller.(ast.MemberExpr).Property, env)

			if err != nil {
				return "", err
			}

			switch v := eval.(type) {
			case interpreter_env.StringVal:
				return v.Value, nil
			default:
				return "", errors.New(compilerErrors.ErrComputedPropertyMustBeString)
			}
		}

		return caller.Caller.(ast.MemberExpr).Property.(ast.Identifier).Symbol, nil
	default:
		return "", nil
	}
}

func EvaluateBodyStmt(body []ast.Stmt, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	var lastEvaluated interpreter_env.RuntimeValue = interpreter_makers.MkNull()

	for _, statement := range body {
		eval, err := Evaluate(statement, env)
		if err != nil {
			return eval, err
		}
		lastEvaluated = eval
	}

	return lastEvaluated, nil
}

/*
 * First return value is the function itself.
 * Second return value is true if the function exists.
 */
func IsNativeFunction(name string) (nativeFns.NativeFunction, bool) {
	function, ok := nativeFns.NativeFunctions[name]

	return function, ok
}
