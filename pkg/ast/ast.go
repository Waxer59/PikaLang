package ast

import "pika/pkg/ast/astTypes"

type Stmt interface {
	GetKind() astTypes.NodeType
}

type Program struct {
	Kind astTypes.NodeType
	Body []Stmt
}

type Identifier struct {
	Kind   astTypes.NodeType
	Symbol string
}

type VariableDeclaration struct {
	Kind       astTypes.NodeType
	Constant   bool
	Identifier string
	Value      Expr
}

type NumericLiteral struct {
	Kind  astTypes.NodeType
	Value int
}

func (p Program) GetKind() astTypes.NodeType {
	return p.Kind
}

func (vd VariableDeclaration) GetKind() astTypes.NodeType {
	return vd.Kind
}

func (i Identifier) GetKind() astTypes.NodeType {
	return i.Kind
}

func (n NumericLiteral) GetKind() astTypes.NodeType {
	return n.Kind
}

func (n NumericLiteral) GetValue() interface{} {
	return n.Value
}
