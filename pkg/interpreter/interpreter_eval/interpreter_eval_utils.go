package interpreter_eval

import (
	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

func EvaluateTruthyFalsyValues(val any) bool {
	switch v := val.(type) {
	case bool:
		return v
	case int, float64, float32:
		return v != 0.0
	case string:
		return v != ""
	case nil:
		return false
	case []ast.Expr:
		return len(v) != 0
	default:
		return true
	}
}

func EvaluateBodyStmt(body []ast.Stmt, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	var lastEvaluated interpreter_env.RuntimeValue = interpreter_makers.MK_Null()

	for _, statement := range body {
		eval, err := Evaluate(statement, env)
		if err != nil {
			return eval, err
		}
		lastEvaluated = eval
	}

	return lastEvaluated, nil
}
