package astTypes

type NodeType string

const (
	// STATEMENTS
	Program             NodeType = "Program"
	VariableDeclaration NodeType = "VariableDeclaration"

	// EXPRESSIONS
	AssigmentExpr NodeType = "AssigmentExpr"
	BinaryExpr    NodeType = "BinaryExpr"
	Identifier    NodeType = "Identifier"

	// LITERALS
	ObjectLiteral  NodeType = "ObjectLiteral"
	Property       NodeType = "Property"
	NumericLiteral NodeType = "NumericLiteral"
	NullLiteral    NodeType = "NullLiteral"
)

var (
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
)
