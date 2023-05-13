package ast

import "pika/pkg/ast/astTypes"

type Stmt interface {
	GetKind() astTypes.NodeType
}

type Program struct {
	Kind astTypes.NodeType
	Body []Stmt
}

func (p Program) GetKind() astTypes.NodeType {
	return p.Kind
}

type Expr interface {
	Stmt
}

type BinaryExpr struct {
	Kind     astTypes.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

func (b BinaryExpr) GetKind() astTypes.NodeType {
	return b.Kind
}

type Identifier struct {
	Kind   astTypes.NodeType
	Symbol string
}

func (i Identifier) GetKind() astTypes.NodeType {
	return i.Kind
}

type NumericLiteral struct {
	Kind  astTypes.NodeType
	Value int
}

func (n NumericLiteral) GetKind() astTypes.NodeType {
	return n.Kind
}

func (n NumericLiteral) GetValue() interface{} {
	return n.Value
}

type NullLiteral struct {
	Kind  astTypes.NodeType
	Value string
}

func (nl NullLiteral) GetKind() astTypes.NodeType {
	return nl.Kind
}
