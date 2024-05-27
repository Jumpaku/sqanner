package types

import (
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

type StructFieldNode interface {
	node.Node
	Named() bool
	Name() node.IdentifierNode
	Type() TypeNode
}

func AcceptStructFieldUnnamed(s *parse.ParseState, fieldType TypeNode) *structFieldNode {
	return &structFieldNode{
		NodeBase:  node.NewNodeBase(s.Begin(), s.End()),
		fieldType: fieldType,
	}
}
func AcceptStructField(s *parse.ParseState, fieldName node.IdentifierNode, fieldType TypeNode) *structFieldNode {
	return &structFieldNode{
		NodeBase:  node.NewNodeBase(s.Begin(), s.End()),
		fieldName: fieldName,
		fieldType: fieldType,
	}
}

func NewStructFieldUnnamed(fieldType TypeNode) *structFieldNode {
	return &structFieldNode{fieldType: fieldType}
}
func NewStructField(fieldName node.IdentifierNode, fieldType TypeNode) *structFieldNode {
	return &structFieldNode{fieldName: fieldName, fieldType: fieldType}
}

type structFieldNode struct {
	node.NodeBase
	fieldName node.IdentifierNode
	fieldType TypeNode
}

func (n *structFieldNode) Children() []node.Node {
	if n.Named() {
		return []node.Node{n.fieldName, n.fieldType}
	}
	return []node.Node{n.fieldType}
}

func (n *structFieldNode) Named() bool {
	return n.fieldName != nil
}

func (n *structFieldNode) Name() node.IdentifierNode {
	return n.fieldName
}

func (n *structFieldNode) Type() TypeNode {
	return n.fieldType
}

var _ StructFieldNode = (*structFieldNode)(nil)
