package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type ArraySyntax int

const (
	ArraySyntaxArray ArraySyntax = iota
	ArraySyntaxTyped
	ArraySyntaxTypeless
)

type ArrayNode interface {
	node.Node
	Syntax() ArraySyntax
	Type() node.TypeNode
	//Values() []ExprNode
}
