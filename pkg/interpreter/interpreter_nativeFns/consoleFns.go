package interpreter_nativeFns

import (
	"fmt"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"

	"github.com/fatih/color"
)

var ConsoleFns = map[string]NativeFunction{
	"print": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		for _, arg := range args {
			fmt.Println(arg.GetValue(), " ")
		}
		return interpreter_makers.MK_NULL()
	},
	"printe": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		for _, arg := range args {
			s := fmt.Sprintf("%v", arg.GetValue())
			color.Red(s)
		}
		return interpreter_makers.MK_NULL()
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
