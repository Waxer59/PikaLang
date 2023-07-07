package interpreter_nativeFns

import (
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

var VarietyFns = map[string]NativeFunction{
	"typeof": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_NaN()
		}

		return interpreter_makers.MK_String(string(args[0].GetType()))
	},
}
