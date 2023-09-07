package ast

import "github.com/Waxer59/PikaLang/pkg/ast/ast_types"

type Property struct {
	Kind  ast_types.NodeType
	Key   string
	Value Expr
}

func (p Property) GetKind() ast_types.NodeType {
	return p.Kind
}

type NumericLiteral struct {
	Kind  ast_types.NodeType
	Value float64
}

func (n NumericLiteral) GetKind() ast_types.NodeType {
	return n.Kind
}

type ObjectLiteral struct {
	Kind       ast_types.NodeType
	Properties []Property
}

func (o ObjectLiteral) GetKind() ast_types.NodeType {
	return o.Kind
}

type NullLiteral struct {
	Kind  ast_types.NodeType
	Value any // nil
}

func (n NullLiteral) GetKind() ast_types.NodeType {
	return n.Kind
}

type BooleanLiteral struct {
	Kind  ast_types.NodeType
	Value bool
}

func (b BooleanLiteral) GetKind() ast_types.NodeType {
	return b.Kind
}

type StringLiteral struct {
	Kind  ast_types.NodeType
	Value string
}

func (s StringLiteral) GetKind() ast_types.NodeType {
	return s.Kind
}

type NaNLiteral struct {
	Kind  ast_types.NodeType
	Value any // nil
}

func (n NaNLiteral) GetKind() ast_types.NodeType {
	return n.Kind
}

type ArrayLiteral struct {
	Kind     ast_types.NodeType
	Elements []Expr
}

func (a ArrayLiteral) GetKind() ast_types.NodeType {
	return a.Kind
}
