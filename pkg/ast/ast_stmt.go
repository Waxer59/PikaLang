package ast

import "pika/pkg/ast/ast_types"

type Stmt interface {
	GetKind() ast_types.NodeType
}

type Program struct {
	Kind ast_types.NodeType
	Body []Stmt
}

func (p Program) GetKind() ast_types.NodeType {
	return p.Kind
}

type FunctionDeclaration struct {
	Kind   ast_types.NodeType
	Params []string
	Name   string
	Body   []Stmt
}

func (f FunctionDeclaration) GetKind() ast_types.NodeType {
	return f.Kind
}

type VariableDeclaration struct {
	Kind       ast_types.NodeType
	Constant   bool
	Identifier string
	Value      Expr
}

func (vd VariableDeclaration) GetKind() ast_types.NodeType {
	return vd.Kind
}
