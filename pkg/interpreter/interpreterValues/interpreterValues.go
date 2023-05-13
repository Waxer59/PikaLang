package interpreterValues

type ValueType string

const (
	Null   ValueType = "null"
	Number ValueType = "number"
)

type RuntimeValue interface {
	GetType() ValueType
}

type NullVal struct {
	Type  ValueType
	Value string
}

func (n NullVal) GetType() ValueType {
	return n.Type
}

type NumberVal struct {
	Type  ValueType
	Value int
}

func (n NumberVal) GetType() ValueType {
	return n.Type
}
