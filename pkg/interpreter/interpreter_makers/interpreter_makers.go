package interpreter_makers

import (
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
)

func MkNull() interpreter_env.NullVal {
	return interpreter_env.NullVal{
		Type:  interpreter_env.Null,
		Value: "null",
	}
}

func MkNumber(n float64) interpreter_env.NumberVal {
	return interpreter_env.NumberVal{
		Type:  interpreter_env.Number,
		Value: n,
	}
}

func MkBoolean(b bool) interpreter_env.BooleanVal {
	return interpreter_env.BooleanVal{
		Type:  interpreter_env.Boolean,
		Value: b,
	}
}

func MkString(s string) interpreter_env.StringVal {
	return interpreter_env.StringVal{
		Type:  interpreter_env.String,
		Value: s,
	}
}

func MkNan() interpreter_env.NaNVal {
	return interpreter_env.NaNVal{
		Type:  interpreter_env.Number,
		Value: "NaN",
	}
}

func MkArray(a []interpreter_env.RuntimeValue) interpreter_env.ArrayVal {
	return interpreter_env.ArrayVal{
		Type:     interpreter_env.Array,
		Elements: a,
	}
}
