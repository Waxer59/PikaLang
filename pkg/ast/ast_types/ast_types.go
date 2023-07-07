package ast_types

type NodeType string

const (
	// STATEMENTS
	Program             NodeType = "Program"
	VariableDeclaration NodeType = "VariableDeclaration"
	FunctionDeclaration NodeType = "FunctionDeclaration"
	IfStatement         NodeType = "IfStatement"
	SwitchStatement     NodeType = "SwitchStatement"

	// EXPRESSIONS
	AssigmentExpr   NodeType = "AssigmentExpr"
	ConditionalExpr NodeType = "ConditionalExp"
	BinaryExpr      NodeType = "BinaryExpr"
	Identifier      NodeType = "Identifier"
	MemberExpr      NodeType = "MemberExpr"
	CallExpr        NodeType = "CallExpr"
	LogicalExpr     NodeType = "LogicalExpr"
	UnaryExpr       NodeType = "UnaryExpr"

	// LITERALS
	ObjectLiteral  NodeType = "ObjectLiteral"
	Property       NodeType = "Property"
	NumericLiteral NodeType = "NumericLiteral"
	NullLiteral    NodeType = "NullLiteral"
	BooleanLiteral NodeType = "BooleanLiteral"
	StringLiteral  NodeType = "StringLiteral"
	NaNLiteral     NodeType = "NaNLiteral"
	ArrayLiteral   NodeType = "ArrayLiteral"
)

var (
	// MATH EXPR
	AdditiveExpr       = []string{"+", "-"}
	MultiplicativeExpr = []string{"*", "/", "%"}
	MathExpr           = []string{"+", "-", "*", "/", "%", "**"}

	// BOOLEAN EXPR
	ComparisonExpr = []string{"<", "<=", ">", ">="}
	EqualityExpr   = []string{"==", "!="}
	BoolExpr       = []string{"<", "<=", ">", ">=", "==", "!="}
)
