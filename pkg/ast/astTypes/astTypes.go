package astTypes

type NodeType string

const (
	Program        NodeType = "Program"
	NumericLiteral NodeType = "NumericLiteral"
	Identifier     NodeType = "Identifier"
	BinaryExpr     NodeType = "BinaryExpr"
)

var (
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
)
