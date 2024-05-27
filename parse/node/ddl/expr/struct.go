package expr

import "github.com/Jumpaku/sqanner/old_parse/node"

type StructSyntax int

const (
	StructSyntaxTuple StructSyntax = iota
	StructSyntaxTyped
	StructSyntaxTypeless
)

type StructNode interface {
	node.Node
	Syntax() StructSyntax
	Type() node.TypeNode
	//FieldValues() []ExprNode
}
