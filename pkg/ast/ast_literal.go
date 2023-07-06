package ast

import "pika/pkg/ast/ast_types"

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

func (n NumericLiteral) GetValue() interface{} {
	return n.Value
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
	Value interface{} // nil
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
	Value interface{}
}

func (n NaNLiteral) GetKind() ast_types.NodeType {
	return n.Kind
}
