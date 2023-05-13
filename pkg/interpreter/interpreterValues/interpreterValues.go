package interpreterValues

type ValueType string

const (
	Null    ValueType = "null"
	Number  ValueType = "number"
	Boolean ValueType = "boolean"
)

type RuntimeValue interface {
	GetType() ValueType
}

type NullVal struct {
	Type  ValueType
	Value interface{} // <-- nil
}

type BooleanVal struct {
	Type  ValueType
	Value bool
}

type NumberVal struct {
	Type  ValueType
	Value int
}

func (n NullVal) GetType() ValueType {
	return n.Type
}

func (b BooleanVal) GetType() ValueType {
	return b.Type
}

func (n NumberVal) GetType() ValueType {
	return n.Type
}
