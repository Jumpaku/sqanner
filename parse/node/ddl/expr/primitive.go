package expr

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
)

type PrimitiveKind int

const (
	PrimitiveKindUnspecified PrimitiveKind = iota
	PrimitiveKindNull
	PrimitiveKindIdentifier
	PrimitiveKindScalar
)

type PrimitiveNode interface {
	node.Node
	PrimitiveKind() PrimitiveKind
	Null() node.KeywordNode
	Identifier() node.IdentifierNode
	Scalar() ScalarNode
}
