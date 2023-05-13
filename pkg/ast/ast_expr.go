package ast

import "pika/pkg/ast/astTypes"

type Expr interface {
	Stmt
}

type AssigmentExpr struct {
	Kind    astTypes.NodeType
	Assigne Expr
	Value   Expr
}

type BinaryExpr struct {
	Kind     astTypes.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

func (a AssigmentExpr) GetKind() astTypes.NodeType {
	return a.Kind
}

func (b BinaryExpr) GetKind() astTypes.NodeType {
	return b.Kind
}
