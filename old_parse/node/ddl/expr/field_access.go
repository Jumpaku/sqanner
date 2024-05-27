package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type FieldAccessNode interface {
	node.Node
	Owner() ExprNode
	Field() node.IdentifierNode
}
