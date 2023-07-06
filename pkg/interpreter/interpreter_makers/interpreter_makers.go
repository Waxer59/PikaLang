package interpreter_makers

import "pika/pkg/interpreter/interpreter_env"

func MK_NULL() interpreter_env.NullVal {
	return interpreter_env.NullVal{
		Type:  interpreter_env.Null,
		Value: nil,
	}
}

func MK_Number(n float64) interpreter_env.NumberVal {
	return interpreter_env.NumberVal{
		Type:  interpreter_env.Number,
		Value: n,
	}
}

func MK_Boolean(b bool) interpreter_env.BooleanVal {
	return interpreter_env.BooleanVal{
		Type:  interpreter_env.Boolean,
		Value: b,
	}
}

func MK_String(s string) interpreter_env.StringVal {
	return interpreter_env.StringVal{
		Type:  interpreter_env.String,
		Value: s,
	}
}

func MK_NaN() interpreter_env.NaNVal {
	return interpreter_env.NaNVal{
		Type:  interpreter_env.Number,
		Value: "NaN",
	}
}
