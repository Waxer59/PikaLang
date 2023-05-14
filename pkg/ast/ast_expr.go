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

func (a AssigmentExpr) GetKind() ast_types.NodeType {
	return a.Kind
}

type BinaryExpr struct {
	Kind     ast_types.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

func (b BinaryExpr) GetKind() ast_types.NodeType {
	return b.Kind
}

type CallExpr struct {
	Kind   ast_types.NodeType
	Args   []Expr
	Caller Expr
}

func (c CallExpr) GetKind() ast_types.NodeType {
	return c.Kind
}

type MemberExpr struct {
	Kind     ast_types.NodeType
	Object   Expr
	Property Expr
	Computed bool
}

func (m MemberExpr) GetKind() ast_types.NodeType {
	return m.Kind
}

type Identifier struct {
	Kind   ast_types.NodeType
	Symbol string
}

func (i Identifier) GetKind() ast_types.NodeType {
	return i.Kind
}
