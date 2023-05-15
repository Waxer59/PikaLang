package interpreter_env

import (
	"fmt"
)

type NativeFunction func(args []RuntimeValue, env Environment) RuntimeValue

var NativeFunctions = map[string]NativeFunction{
	"print": func(args []RuntimeValue, env Environment) RuntimeValue {
		for _, arg := range args {
			fmt.Print(arg)
		}
		return NullVal{
			Type:  Null,
			Value: nil,
		}
	},
}
