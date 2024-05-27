package types

import (
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

type TypeSizeNode interface {
	node.Node
	Max() bool
	Size() int
}

func AcceptTypeSize(s *parse.ParseState, size int) *typeSizeNode {
	return &typeSizeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		size:     size,
	}
}

func AcceptTypeSizeMax(s *parse.ParseState) *typeSizeNode {
	return &typeSizeNode{
		NodeBase: node.NewNodeBase(s.Begin(), s.End()),
		max:      true,
	}
}

func NewTypeSize(size int) *typeSizeNode {
	return &typeSizeNode{size: size}
}

func NewTypeSizeMax() *typeSizeNode {
	return &typeSizeNode{max: true}
}

type typeSizeNode struct {
	node.NodeBase
	size int
	max  bool
}

var _ TypeSizeNode = (*typeSizeNode)(nil)

func (n *typeSizeNode) Children() []node.Node {
	return nil
}

func (n *typeSizeNode) Max() bool {
	return n.max
}

func (n *typeSizeNode) Size() int {
	return n.size
}
