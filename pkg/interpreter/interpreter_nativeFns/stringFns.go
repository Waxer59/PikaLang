package interpreter_nativeFns

import (
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
	"strings"
)

var StringFns = map[string]NativeFunction{
	"toUpperCase": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.String {
			return interpreter_makers.MK_String("")
		}

		str := args[0].GetValue().(string)
		result := strings.ToUpper(str)
		return interpreter_makers.MK_String(result)
	},
	"toLowerCase": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.String {
			return interpreter_makers.MK_String("")
		}

		str := args[0].GetValue().(string)
		result := strings.ToLower(str)
		return interpreter_makers.MK_String(result)
	},
	"capitalize": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.String {
			return interpreter_makers.MK_String("")
		}

		str := args[0].GetValue().(string)
		if len(str) <= 0 {
			return interpreter_makers.MK_String("")
		}
		result := strings.ToUpper(str[:1]) + str[1:]
		return interpreter_makers.MK_String(result)
	},
	"startsWith": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 2 || args[0].GetType() != interpreter_env.String || args[1].GetType() != interpreter_env.String {
			return interpreter_makers.MK_Boolean(false)
		}

		str := args[0].GetValue().(string)
		prefix := args[1].GetValue().(string)
		result := strings.HasPrefix(str, prefix)
		return interpreter_makers.MK_Boolean(result)
	},
	"endsWith": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 2 || args[0].GetType() != interpreter_env.String || args[1].GetType() != interpreter_env.String {
			return interpreter_makers.MK_Boolean(false)
		}

		str := args[0].GetValue().(string)
		suffix := args[1].GetValue().(string)
		result := strings.HasSuffix(str, suffix)
		return interpreter_makers.MK_Boolean(result)
	},
	"reverseString": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) < 1 || args[0].GetType() != interpreter_env.String {
			return interpreter_makers.MK_String("")
		}

		str := args[0].GetValue().(string)
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result := string(runes)
		return interpreter_makers.MK_String(result)
	},
}
