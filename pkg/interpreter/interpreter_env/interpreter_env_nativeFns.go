package interpreter_env

import (
	"fmt"
)

type NativeFunction func(args []RuntimeValue, env Environment) RuntimeValue

var NativeFunctions = map[string]NativeFunction{
	"print": func(args []RuntimeValue, env Environment) RuntimeValue {
		for _, arg := range args {
			fmt.Println(arg.GetValue(), " ")
		}
		return NullVal{
			Type:  Null,
			Value: nil,
		}
	},
	"len": func(args []RuntimeValue, env Environment) RuntimeValue {
		arg, ok := args[0].GetValue().(string)
		if !ok {
			return NullVal{
				Type:  Null,
				Value: nil,
			}
		}
		return NumberVal{
			Type:  Number,
			Value: float64(len(arg)),
		}
	},
}
