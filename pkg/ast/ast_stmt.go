package ast

import (
	"pika/pkg/ast/ast_types"
)

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

type IfStatement struct {
	Kind       ast_types.NodeType
	Test       Expr
	Body       []Stmt
	ElseIfStmt []ElseIfStatement
	ElseBody   []Stmt
}

type ElseIfStatement struct {
	Test Expr
	Body []Stmt
}

func (cd IfStatement) GetKind() ast_types.NodeType {
	return cd.Kind
}

type SwitchStatement struct {
	Kind         ast_types.NodeType
	Discriminant Expr
	CaseStmts    []CaseStatement
	DefaultStmt  CaseStatement
}

type CaseStatement struct {
	Test []Expr
	Body []Stmt
}

func (cs SwitchStatement) GetKind() ast_types.NodeType {
	return cs.Kind
}

type ReturnStatement struct {
	Kind     ast_types.NodeType
	Argument Expr
}

func (rs ReturnStatement) GetKind() ast_types.NodeType {
	return rs.Kind
}

type WhileStatement struct {
	Kind ast_types.NodeType
	Test Expr
	Body []Stmt
}

func (ws WhileStatement) GetKind() ast_types.NodeType {
	return ws.Kind
}
