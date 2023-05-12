package ast

import "pikalang/pkg/ast/astTypes"

type Stmt struct {
	Kind astTypes.NodeType
}

type Program struct {
	Kind astTypes.NodeType
	Body []interface{}
}

type Expr interface{}

type BinaryExpr struct {
	Kind     astTypes.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

type Identifier struct {
	Kind   astTypes.NodeType
	Symbol string
}

type NumericLiteral struct {
	Kind  astTypes.NodeType
	Value int
}
