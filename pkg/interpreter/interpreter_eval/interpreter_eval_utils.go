package interpreter_eval

func EvaluateTruthyFalsyValues(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case int, float64, float32:
		return v != 0.0
	case string:
		return v != ""
	case nil:
		return false
	default:
		return true
	}
}
