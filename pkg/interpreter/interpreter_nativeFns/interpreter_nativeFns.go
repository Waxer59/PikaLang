package interpreter_nativeFns

import (
	"pika/internal/utils"
	"pika/pkg/interpreter/interpreter_env"
)

type NativeFunction func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue

var NativeFunctions = utils.MergeMaps(BooleanFns, ConsoleFns, NumberFns, ParseFns, StringFns, VarietyFns)
