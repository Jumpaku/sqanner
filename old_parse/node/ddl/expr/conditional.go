package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type ConditionalBranchNode interface {
	node.Node
	IsElse() bool
	When() ExprNode
	Then() ExprNode
	Else() ExprNode
}
type ConditionalNode interface {
	node.Node
	CaseExpr() ExprNode
	Branches() []ConditionalBranchNode
	Type() node.TypeNode
}
