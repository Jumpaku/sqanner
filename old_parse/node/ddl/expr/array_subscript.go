package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type ArraySubscriptSpecifierKind int

const (
	ArraySubscriptSpecifierKindUnspecified ArraySubscriptSpecifierKind = iota
	ArraySubscriptSpecifierKindOffset
	ArraySubscriptSpecifierKindSafeOffset
	ArraySubscriptSpecifierKindOrdinal
	ArraySubscriptSpecifierKindSafeOrdinal
)

type ArraySubscriptSpecifierNode interface {
	node.Node
	SpecifierKind() ArraySubscriptSpecifierKind
	Index() ExprNode
}

type ArraySubscriptNode interface {
	node.Node
	Array() ExprNode
	Specifier() ArraySubscriptSpecifierNode
}
