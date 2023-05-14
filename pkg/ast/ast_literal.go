package ast

import "pika/pkg/ast/ast_types"

type NumericLiteral struct {
	Kind  ast_types.NodeType
	Value int
}

type Property struct {
	Kind  ast_types.NodeType
	Key   string
	Value Expr
}

type ObjectLiteral struct {
	Kind       ast_types.NodeType
	Properties []Property
}

func (o ObjectLiteral) GetKind() ast_types.NodeType {
	return o.Kind
}

func (p Property) GetKind() ast_types.NodeType {
	return p.Kind
}

func (n NumericLiteral) GetKind() ast_types.NodeType {
	return n.Kind
}

func (n NumericLiteral) GetValue() interface{} {
	return n.Value
}
