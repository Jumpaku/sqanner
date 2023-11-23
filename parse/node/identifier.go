package node

import (
	"github.com/Jumpaku/sqanner/parse"
	"strings"
)

type IdentifierNode interface {
	Node
	isIdentifierNode()
	Value() string
	IsQuoted() bool
	UnquotedValue() string
}

var _ IdentifierNode = (*identifierNode)(nil)

type identifierNode struct {
	NodeBase
	value string
}

func AcceptIdentifier(s *parse.ParseState, value string) *identifierNode {
	return &identifierNode{
		NodeBase: NewNodeBase(s.Begin(), s.End()),
		value:    value,
	}
}

func NewIdentifier(value string) *identifierNode {
	return &identifierNode{value: value}
}

func (n *identifierNode) Value() string {
	return n.value
}

func (n *identifierNode) IsQuoted() bool {
	return strings.HasPrefix(n.Value(), "`")
}

func (n *identifierNode) UnquotedValue() string {
	if n.IsQuoted() {
		return n.Value()
	}
	return n.value[1 : len(n.value)-1]
}

func (n *identifierNode) Children() []Node {
	return nil
}

func (n *identifierNode) isIdentifierNode() {}
