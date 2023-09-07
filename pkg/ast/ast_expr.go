package ast

import "github.com/Waxer59/PikaLang/pkg/ast/ast_types"

type Expr interface {
	Stmt
}

type AssigmentExpr struct {
	Kind     ast_types.NodeType
	Assigne  Expr
	Value    Expr
	Operator string
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

func (c CallExpr) GetFnName() string {
	switch c.Caller.(type) {
	case Identifier:
		return c.Caller.(Identifier).Symbol
	case MemberExpr:
		return c.Caller.(MemberExpr).Property.(Identifier).Symbol
	default:
		return ""
	}
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

type ConditionalExpr struct {
	Kind       ast_types.NodeType
	Condition  Expr
	Consequent Expr
	Alternate  Expr
}

func (c ConditionalExpr) GetKind() ast_types.NodeType {
	return c.Kind
}

type LogicalExpr struct {
	Kind     ast_types.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

func (l LogicalExpr) GetKind() ast_types.NodeType {
	return l.Kind
}

type UnaryExpr struct {
	Kind     ast_types.NodeType
	Operator string
	Argument Expr
	Prefix   bool
}

func (u UnaryExpr) GetKind() ast_types.NodeType {
	return u.Kind
}

type UpdateExpr struct {
	Kind     ast_types.NodeType
	Operator string
	Argument Identifier
	Prefix   bool
}

func (u UpdateExpr) GetKind() ast_types.NodeType {
	return u.Kind
}

type ArrowFunctionExpr struct {
	Kind   ast_types.NodeType
	Params []Identifier
	Body   []Stmt
}

func (a ArrowFunctionExpr) GetKind() ast_types.NodeType {
	return a.Kind
}
