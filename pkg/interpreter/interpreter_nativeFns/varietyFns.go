package interpreter_nativeFns

import (
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

var VarietyFns = map[string]NativeFunction{
	"len": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_NaN()
		}

		switch args[0].GetType() {
		case interpreter_env.String:
			arg, ok := args[0].GetValue().(string)
			if !ok {
				return interpreter_makers.MK_NaN()
			}
			return interpreter_makers.MK_Number(float64(len(arg)))
		case interpreter_env.Array:
			arg, ok := args[0].GetValue().([]interpreter_env.RuntimeValue)
			if !ok {
				return interpreter_makers.MK_NaN()
			}
			return interpreter_makers.MK_Number(float64(len(arg)))
		default:
			return interpreter_makers.MK_NaN()
		}

	},
	"typeof": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_NaN()
		}

		return interpreter_makers.MK_String(string(args[0].GetType()))
	},
}
