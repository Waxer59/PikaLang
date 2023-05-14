package ast_types

type NodeType string

const (
	// STATEMENTS
	Program             NodeType = "Program"
	VariableDeclaration NodeType = "VariableDeclaration"
	FunctionDeclaration NodeType = "FunctionDeclaration"

	// EXPRESSIONS
	AssigmentExpr NodeType = "AssigmentExpr"
	BinaryExpr    NodeType = "BinaryExpr"
	Identifier    NodeType = "Identifier"
	MemberExpr    NodeType = "MemberExpr"
	CallExpr      NodeType = "CallExpr"

	// LITERALS
	ObjectLiteral  NodeType = "ObjectLiteral"
	Property       NodeType = "Property"
	NumericLiteral NodeType = "NumericLiteral"
	NullLiteral    NodeType = "NullLiteral"
	BooleanLiteral NodeType = "BooleanLiteral"
	StringLiteral  NodeType = "StringLiteral"
)

var (
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
)
