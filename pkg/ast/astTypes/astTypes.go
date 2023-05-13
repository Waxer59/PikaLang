package astTypes

type NodeType string

const (
	Program        NodeType = "Program"
	NumericLiteral NodeType = "NumericLiteral"
	Identifier     NodeType = "Identifier"
	BinaryExpr     NodeType = "BinaryExpr"
	NullLiteral    NodeType = "NullLiteral"
)

var (
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
)
