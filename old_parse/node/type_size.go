package node

type TypeSizeNode interface {
	Node
	Max() bool
	Size() int
}

func TypeSize(size int) NewNodeFunc[TypeSizeNode] {
	return func(begin int, end int) TypeSizeNode {
		return typeSize{
			NodeBase: NodeBase{kind: KindTypeSize, begin: begin, end: end},
			size:     size,
		}
	}
}

func TypeSizeMax() NewNodeFunc[TypeSizeNode] {
	return func(begin int, end int) TypeSizeNode {
		return typeSize{
			NodeBase: NodeBase{kind: KindTypeSize, begin: begin, end: end},
			max:      true,
		}
	}
}

type typeSize struct {
	NodeBase
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
