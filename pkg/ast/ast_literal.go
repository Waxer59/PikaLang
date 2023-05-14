package ast

import "pika/pkg/ast/astTypes"

type NumericLiteral struct {
	Kind  astTypes.NodeType
	Value int
}

type Property struct {
	Kind  astTypes.NodeType
	Key   string
	Value Expr
}

type ObjectLiteral struct {
	Kind       astTypes.NodeType
	Properties []Property
}

func (o ObjectLiteral) GetKind() astTypes.NodeType {
	return o.Kind
}

func (p Property) GetKind() astTypes.NodeType {
	return p.Kind
}

func (n NumericLiteral) GetKind() astTypes.NodeType {
	return n.Kind
}

func (n NumericLiteral) GetValue() interface{} {
	return n.Value
}
