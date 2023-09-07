package interpreter_nativeFns

import (
	"fmt"

	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

var ConsoleFns = map[string]NativeFunction{
	"print": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {

		for _, arg := range args {
			printPrimitive(arg)
			fmt.Println("")
		}

		return interpreter_makers.MK_Null()
	},
	"prompt": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 {
			return interpreter_makers.MK_String("")
		}

		fmt.Print(args[0].GetValue())
		var input string
		fmt.Scanln(&input)
		return interpreter_makers.MK_String(input)
	},
}

func printPrimitive(val interpreter_env.RuntimeValue) {
	switch val.GetType() {
	case interpreter_env.Array:
		arr, _ := val.GetValue().([]interpreter_env.RuntimeValue)
		fmt.Print("[ ")
		for idx, el := range arr {
			printPrimitive(el)
			if idx != len(arr)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print(" ]")
	case interpreter_env.Object:
		obj, _ := val.GetValue().(map[string]interpreter_env.RuntimeValue)
		fmt.Print("{ ")
		for key, value := range obj {
			fmt.Print(key + ": ")
			printPrimitive(value)
			fmt.Print(", ")
		}
		fmt.Print("}")
	case interpreter_env.String:
		fmt.Print("\"" + val.GetValue().(string) + "\"")
	case interpreter_env.Function, interpreter_env.ArrowFunction:
		fn, _ := val.GetValue().(interpreter_env.FunctionVal)
		if fn.Name != nil {
			fmt.Print("fn " + *fn.Name + " ")
		}
		fmt.Print("(")
		for idx, arg := range fn.Params {
			fmt.Println(arg.Symbol)
			if idx != len(fn.Params)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print(") { ... }")
	default:
		fmt.Print(val.GetValue())
	}
}
