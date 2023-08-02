package node

type NewNodeFunc[T Node] func(begin int, end int) T

type Node interface {
	Kind() NodeKind
	Begin() int
	End() int
	Children() []Node
}

type nodeBase struct {
	begin int
	end   int
	kind  NodeKind
}

func (n nodeBase) Kind() NodeKind {
	return n.kind
}

func (n nodeBase) Begin() int {
	return n.begin
}

func (n nodeBase) End() int {
	return n.end
}
