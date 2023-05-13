package astTypes

type NodeType string

const (
	// STATEMENTS
	Program             NodeType = "Program"
	VariableDeclaration NodeType = "VariableDeclaration"

	// EXPRESSIONS
	NumericLiteral NodeType = "NumericLiteral"
	Identifier     NodeType = "Identifier"
	BinaryExpr     NodeType = "BinaryExpr"
	NullLiteral    NodeType = "NullLiteral"
)

var (
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
)
