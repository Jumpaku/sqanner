package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type FunctionCallNode interface {
	node.Node
	Func() ExprNode
	Args() []ExprNode
}
