package node

type Node interface {
	Begin() int
	End() int
	Len() int
	Children() []Node
}

type NodeBase struct {
	begin int
	end   int
}

func NewNodeBase(begin, end int) NodeBase { return NodeBase{begin: begin, end: end} }

func (n NodeBase) Begin() int {
	return n.begin
}

func (n NodeBase) End() int {
	return n.end
}

func (n NodeBase) Len() int {
	return n.end - n.begin
}
