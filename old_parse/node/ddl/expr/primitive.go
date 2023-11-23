package expr

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
)

type PrimitiveKind int

const (
	PrimitiveKindUnspecified PrimitiveKind = iota
	PrimitiveKindNull
	PrimitiveKindIdentifier
	PrimitiveKindLiteral
)

type PrimitiveNode interface {
	node.Node
	PrimitiveKind() PrimitiveKind
	Null() node.KeywordNode
	Identifier() node.IdentifierNode
	LiteralType() node.TypeCode
	Int64Literal() Int64LiteralNode
	Float64Literal() Float64LiteralNode
	BoolLiteral() BoolLiteralNode
	StringLiteral() StringLiteralNode
	BytesLiteral() StringLiteralNode
	TimestampLiteral() TimestampLiteralNode
	DateLiteral() DateLiteralNode
	JSONLiteral() JSONLiteralNode
	NumericLiteral() NumericLiteralNode
	StructLiteral() StructLiteralNode
	ArrayLiteral() ArrayLiteralNode
}
