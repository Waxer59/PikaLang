package interpreter_makers

import "pika/pkg/interpreter/interpreter_env"

func MK_NULL() interpreter_env.NullVal {
	return interpreter_env.NullVal{
		Type:  interpreter_env.Null,
		Value: nil,
	}
}

func MK_Number(n int) interpreter_env.NumberVal {
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
