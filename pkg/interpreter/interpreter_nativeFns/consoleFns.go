package interpreter_nativeFns

import (
	"fmt"

	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

var ConsoleFns = map[string]NativeFunction{
	"print": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {

		for _, arg := range args {
			switch arg.GetType() {
			case interpreter_env.Array:
				arr, _ := arg.GetValue().([]interpreter_env.RuntimeValue)
				fmt.Print("[ ")
				for idx, el := range arr {
					fmt.Print(el.GetValue())
					if idx != len(arr)-1 {
						fmt.Print(", ")
					}
				}
				fmt.Println(" ]")
			case interpreter_env.Object:
				obj, _ := arg.GetValue().(map[string]interpreter_env.RuntimeValue)
				objLen := len(obj) - 1
				idx := 0

				fmt.Print("{ ")
				for key, value := range obj {
					fmt.Print(key, ": ", value.GetValue())

					idx++
					if idx < objLen {
						fmt.Print(", ")
					}
					fmt.Print(", ")
				}
				fmt.Println("}")
			default:
				fmt.Println(arg.GetValue(), " ")
			}
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
