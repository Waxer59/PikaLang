package nativeFns

import (
	"github.com/Waxer59/PikaLang/internal/utils"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
)

type NativeFunction func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue

var NativeFunctions = utils.MergeMaps(BooleanFns, ConsoleFns, NumberFns, ParseFns, StringFns, VarietyFns, ArrayFns)
