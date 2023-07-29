package node

import (
	"github.com/Jumpaku/sqanner/tokenize"
)

type PathNode interface {
	Node
	Identifiers() []IdentifierNode
}

func Path(identifiers []IdentifierNode) nodeFunc[PathNode] {
	return func(head int, tokens []tokenize.Token) PathNode {
		return path{
			nodeBase:    nodeBase{kind: NodePath, begin: head, tokens: tokens},
			identifiers: identifiers,
		}
	}
}

type path struct {
	nodeBase
	identifiers []IdentifierNode
}

var _ Node = path{}
var _ PathNode = path{}

func (n path) Children() []Node {
	ch := []Node{}
	for _, identifier := range n.identifiers {
		ch = append(ch, identifier)
	}
	return ch
}

func (n path) Identifiers() []IdentifierNode {
	return n.identifiers
}
