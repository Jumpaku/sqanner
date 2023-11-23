package node

type PathNode interface {
	Node
	Identifiers() []IdentifierNode
}

func Path(identifiers []IdentifierNode) NewNodeFunc[PathNode] {
	return func(begin, end int) PathNode {
		return path{
			NodeBase:    NodeBase{kind: KindPath, begin: begin, end: end},
			identifiers: identifiers,
		}
	}
}

type path struct {
	NodeBase
	identifiers []IdentifierNode
}

var _ Node = path{}
var _ PathNode = path{}

func (n path) Children() []Node {
	children := []Node{}
	for _, identifier := range n.identifiers {
		children = append(children, identifier)
	}
	return children
}

func (n path) Identifiers() []IdentifierNode {
	return n.identifiers
}
