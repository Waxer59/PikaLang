package interpreter_nativeFns

import (
	"math"
	"math/rand"
	"time"

	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_makers"
)

var NumberFns = map[string]NativeFunction{
	"randNum": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) <= 1 {
			return interpreter_makers.MK_NaN()
		}

		min := int(args[0].GetValue().(float64))
		max := int(args[1].GetValue().(float64))

		if min > max {
			return interpreter_makers.MK_NaN()
		}
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(max-min+1) + min
		return interpreter_makers.MK_Number(float64(num))
	},
	"pow": func(args []interpreter_env.RuntimeValue, env interpreter_env.Environment) interpreter_env.RuntimeValue {
		if len(args) <= 1 {
			return interpreter_makers.MK_NaN()
		}

		base := args[0].GetValue().(float64)
		exponent := args[1].GetValue().(float64)
		result := math.Pow(base, exponent)
		return interpreter_makers.MK_Number(result)
	},
}
