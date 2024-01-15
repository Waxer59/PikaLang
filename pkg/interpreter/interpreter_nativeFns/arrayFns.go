package interpreter_nativeFns

import (
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
	"golang.org/x/exp/slices"
)

var ArrayFns = map[string]NativeFunction{
	"includes": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 2 || args[0].GetType() != interpreter_env.Array {
			return interpreter_makers.MkBoolean(false)
		}

		return interpreter_makers.MkBoolean(slices.Contains(args[0].(interpreter_env.ArrayVal).GetElements(), args[1].GetValue()))
	},
	"push": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 2 || args[0].GetType() != interpreter_env.Array {
			return interpreter_makers.MkNull()
		}

		arr := args[0].GetValue().([]interpreter_env.RuntimeValue)

		arr = append(arr, args[1:]...)

		return interpreter_makers.MkArray(arr)
	},
	"pop": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.Array {
			return interpreter_makers.MkNull()
		}
		arr := args[0].GetValue().([]interpreter_env.RuntimeValue)
		arr = arr[:len(arr)-1]

		return interpreter_makers.MkArray(arr)
	},
	"shift": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.Array {
			return interpreter_makers.MkNull()
		}
		arr := args[0].GetValue().([]interpreter_env.RuntimeValue)
		arr = arr[len(arr)-1:]

		return interpreter_makers.MkArray(arr)
	},
	"indexOf": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 2 || args[0].GetType() != interpreter_env.Array {
			return interpreter_makers.MkNull()
		}

		arr := args[0].(interpreter_env.ArrayVal).GetElements()
		searchElement := args[1].GetValue()

		for index, element := range arr {
			if element == searchElement {
				return interpreter_makers.MkNumber(float64(index))
			}
		}

		return interpreter_makers.MkNumber(-1)
	},
}
