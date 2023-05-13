package interpreterValues

type ValueType string

const (
	Null   ValueType = "null"
	Number ValueType = "number"
)

type RuntimeValue interface{}

type NullVal struct {
	Type  string
	Value string
}

type NumberVal struct {
	Type  string
	Value int
}
