package ast

import "pika/pkg/ast/ast_types"

type Expr interface {
	Stmt
}

type AssigmentExpr struct {
	Kind    ast_types.NodeType
	Assigne Expr
	Value   Expr
}

type BinaryExpr struct {
	Kind     ast_types.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

type CallExpr struct {
	Kind   ast_types.NodeType
	Args   []Expr
	Caller Expr
}

type MemberExpr struct {
	Kind     ast_types.NodeType
	Object   Expr
	Property Expr
	Computed bool
}

type Identifier struct {
	Kind   ast_types.NodeType
	Symbol string
}

func (i Identifier) GetKind() ast_types.NodeType {
	return i.Kind
}

func (m MemberExpr) GetKind() ast_types.NodeType {
	return m.Kind
}

func (c CallExpr) GetKind() ast_types.NodeType {
	return c.Kind
}

func (a AssigmentExpr) GetKind() ast_types.NodeType {
	return a.Kind
}

func (b BinaryExpr) GetKind() ast_types.NodeType {
	return b.Kind
}
