package ast

import "pika/pkg/ast/astTypes"

type VariableDeclaration struct {
	Kind       astTypes.NodeType
	Constant   bool
	Identifier string
	Value      Expr
}

type Program struct {
	Kind astTypes.NodeType
	Body []Stmt
}

func (p Program) GetKind() astTypes.NodeType {
	return p.Kind
}

func (vd VariableDeclaration) GetKind() astTypes.NodeType {
	return vd.Kind
}
