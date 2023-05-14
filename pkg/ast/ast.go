package ast

import "pika/pkg/ast/astTypes"

type Stmt interface {
	GetKind() astTypes.NodeType
}
