package node

type NewNodeFunc[T Node] func(begin int, end int) T

type Node interface {
	Kind() Kind
	Begin() int
	End() int
	Len() int
	Children() []Node
}

type nodeBase struct {
	begin int
	end   int
	kind  Kind
}

func (n nodeBase) Kind() Kind {
	return n.kind
}

func (n nodeBase) Begin() int {
	return n.begin
}

func (n nodeBase) End() int {
	return n.end
}

func (n nodeBase) Len() int {
	return n.end - n.begin
}
