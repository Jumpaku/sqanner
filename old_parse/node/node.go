package node

type NewNodeFunc[T Node] func(begin int, end int) T

type Node interface {
	Kind() Kind
	Begin() int
	End() int
	Len() int
	Children() []Node
}

type NodeBase struct {
	begin int
	end   int
	kind  Kind
}

func (n NodeBase) Kind() Kind {
	return n.kind
}

func (n NodeBase) Begin() int {
	return n.begin
}

func (n NodeBase) End() int {
	return n.end
}

func (n NodeBase) Len() int {
	return n.end - n.begin
}
