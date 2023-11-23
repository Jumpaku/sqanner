package expr

import (
	"cloud.google.com/go/civil"
	"github.com/Jumpaku/sqanner/old_parse/node"
	"math/big"
	"time"
)

type Int64LiteralNode interface {
	node.Node
	Value() int64
}

type Float64LiteralNode interface {
	node.Node
	Value() float64
}
type BoolLiteralNode interface {
	node.Node
	Value() bool
}
type StringLiteralNode interface {
	node.Node
	Value() string
}
type BytesLiteralNode interface {
	node.Node
	Value() []byte
	StringLiteral() StringLiteralNode
}
type TimestampLiteralNode interface {
	node.Node
	Value() time.Time
	StringLiteral() StringLiteralNode
}
type DateLiteralNode interface {
	node.Node
	Value() civil.Date
	StringLiteral() StringLiteralNode
}
type JSONLiteralNode interface {
	node.Node
	BytesValue() []byte
	StringLiteral() StringLiteralNode
}
type NumericLiteralNode interface {
	node.Node
	Value() big.Rat
	StringLiteral() StringLiteralNode
}

type ArrayLiteralSyntax int

const (
	ArrayLiteralSyntaxArray ArrayLiteralSyntax = iota
	ArrayLiteralSyntaxTyped
	ArrayLiteralSyntaxTypeless
)

type ArrayLiteralNode interface {
	node.Node
	Syntax() ArrayLiteralSyntax
	Type() node.TypeNode
	Values() []ExprNode
}

type StructLiteralSyntax int

const (
	StructLiteralSyntaxTuple StructLiteralSyntax = iota
	StructLiteralSyntaxTyped
	StructLiteralSyntaxTypeless
)

type StructLiteralNode interface {
	node.Node
	Syntax() StructLiteralSyntax
	Type() node.TypeNode
	FieldValues() []ExprNode
}
