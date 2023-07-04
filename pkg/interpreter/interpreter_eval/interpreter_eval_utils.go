package interpreter_eval

import (
	"errors"
	compilerErrors "pika/internal/errors"
)

func EvaluateTruthyFalsyValues(val interface{}) (bool, error) {
	boolVal, ok := val.(bool)

	if !ok {
		return false, errors.New(compilerErrors.ErrTypesInvalidType)
	}

	switch val.(type) {
	case string:
		val = val.(string) != ""
	case int, float32, float64:
		val = val != 0.0
	default:
		val = false
	}

	return boolVal, nil
}
