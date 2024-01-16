package nativeFns

import (
	"fmt"
	"strconv"

	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

func EvaluateTruthyFalsyValues(runtime interpreter_env.RuntimeValue) bool {
	val := runtime.GetValue()
	result := true
	switch runtime.GetType() {
	case interpreter_env.Boolean:
		result = val.(bool)
	case interpreter_env.Null:
		result = false
	case interpreter_env.Number:
		result = val != 0
	case interpreter_env.Array:
		result = len(val.([]interpreter_env.RuntimeValue)) > 0
	case interpreter_env.String:
		result = len(val.(string)) > 0
	}

	return result
}

var ParseFns = map[string]NativeFunction{
	"string": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MkString("")
		}

		switch args[0].GetType() {
		case interpreter_env.Null:
			return interpreter_makers.MkString("null")
		case interpreter_env.Object:
			return interpreter_makers.MkString("object")
		case interpreter_env.Array:
			arr := args[0].GetValue().([]interpreter_env.RuntimeValue)
			s := "["
			for i, v := range arr {
				s += fmt.Sprintf("%v", v.GetValue())
				if i != len(arr)-1 {
					s += ", "
				}
			}
			s += "]"
			return interpreter_makers.MkString(s)
		default:
			s := fmt.Sprintf("%v", args[0].GetValue())
			return interpreter_makers.MkString(s)
		}
	},
	"num": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MkNan()
		}

		i, err := strconv.ParseFloat(args[0].GetValue().(string), 64)

		if err != nil {
			return interpreter_makers.MkNan()
		}

		return interpreter_makers.MkNumber(i)
	},
	"bool": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MkBoolean(false)
		}

		result := EvaluateTruthyFalsyValues(args[0])

		return interpreter_makers.MkBoolean(result)
	},
}
