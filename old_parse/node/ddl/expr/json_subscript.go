package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type JSONSubscriptNode interface {
	node.Node
	Owner() ExprNode
	Subscript() ExprNode
}
