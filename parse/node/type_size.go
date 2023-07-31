package node

import "github.com/Jumpaku/sqanner/tokenize"

type TypeSizeNode interface {
	Node
	Max() bool
	Size() int
}

func TypeSize(size int) NewNodeFunc[TypeSizeNode] {
	return func(begin int, tokens []tokenize.Token) TypeSizeNode {
		return typeSize{
			nodeBase: nodeBase{kind: NodeTypeSize, begin: begin, tokens: tokens},
			size:     size,
		}
	}
}

func TypeSizeMax() NewNodeFunc[TypeSizeNode] {
	return func(begin int, tokens []tokenize.Token) TypeSizeNode {
		return typeSize{
			nodeBase: nodeBase{kind: NodeTypeSize, begin: begin, tokens: tokens},
			max:      true,
		}
	}
}

type typeSize struct {
	nodeBase
	size int
	max  bool
}

var _ TypeSizeNode = typeSize{}

func (n typeSize) Children() []Node {
	return nil
}

func (n typeSize) Max() bool {
	return n.max
}

func (n typeSize) Size() int {
	return n.size
}
