package interpreter

func MK_NULL() NullVal {
	return NullVal{
		Type:  Null,
		Value: nil,
	}
}

func MK_Number(n int) NumberVal {
	return NumberVal{
		Type:  Number,
		Value: n,
	}
}

func MK_Boolean(b bool) BooleanVal {
	return BooleanVal{
		Type:  Boolean,
		Value: b,
	}
}

func MK_NATIVE_FN(call FunctionCall) NativeFnVal {
	return NativeFnVal{
		Type: NativeFn,
		Call: call,
	}
}
