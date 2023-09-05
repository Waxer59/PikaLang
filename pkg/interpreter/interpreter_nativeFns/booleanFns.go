package interpreter_nativeFns

import (
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

var BooleanFns = map[string]NativeFunction{
	"isNaN": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_Boolean(true)
		}

		return interpreter_makers.MK_Boolean(args[0].GetValue() == "NaN" && args[0].GetType() == interpreter_env.Number)
	},
	"isNull": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_Boolean(false)
		}

		return interpreter_makers.MK_Boolean(args[0].GetValue() == nil)
	},
}
