package interpreterMakers

import "pika/pkg/interpreter/interpreterValues"

func MK_NULL() interpreterValues.NullVal {
	return interpreterValues.NullVal{
		Type:  interpreterValues.Null,
		Value: nil,
	}
}

func MK_Number(n int) interpreterValues.NumberVal {
	return interpreterValues.NumberVal{
		Type:  interpreterValues.Number,
		Value: n,
	}
}

func MK_Boolean(b bool) interpreterValues.BooleanVal {
	return interpreterValues.BooleanVal{
		Type:  interpreterValues.Boolean,
		Value: b,
	}
}
