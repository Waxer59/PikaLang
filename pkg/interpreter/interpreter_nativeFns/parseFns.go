package interpreter_nativeFns

import (
	"fmt"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
	"strconv"
)

var ParseFns = map[string]NativeFunction{
	"string": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_String("")
		}

		switch args[0].GetType() {
		case interpreter_env.Null:
			return interpreter_makers.MK_String("null")
		case interpreter_env.Object:
			return interpreter_makers.MK_String("object")
		default:
			s := fmt.Sprintf("%v", args[0].GetValue())
			return interpreter_makers.MK_String(s)
		}
	},
	"num": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_NaN()
		}

		i, err := strconv.ParseFloat(args[0].GetValue().(string), 64)

		if err != nil {
			return interpreter_makers.MK_NaN()
		}

		return interpreter_makers.MK_Number(i)
	},
	"bool": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_Boolean(false)
		}

		result := false
		switch v := args[0].GetValue().(type) {
		case bool:
			result = v
		case int, float64, float32:
			result = v != 0.0
		case string:
			result = v != ""
		case nil:
			result = false
		default:
			result = true
		}

		return interpreter_makers.MK_Boolean(result)
	},
}
