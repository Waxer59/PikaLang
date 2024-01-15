package interpreter_nativeFns

import (
	"fmt"
	"strconv"

	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

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

		return interpreter_makers.MkBoolean(result)
	},
}
