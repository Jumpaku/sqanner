package node

import (
	"github.com/Jumpaku/sqanner/parse"
	"strings"
)

type identifierNode interface {
	Value() string
	IsQuoted() bool
	UnquotedValue() string
}

var _ Node = (*IdentifierNode)(nil)
var _ identifierNode = (*IdentifierNode)(nil)

type IdentifierNode struct {
	NodeBase
	value string
}

func AcceptIdentifier(s *parse.ParseState, value string) *IdentifierNode {
	return &IdentifierNode{
		NodeBase: NewNodeBase(s.Begin(), s.End()),
		value:    value,
	}
}

func NewIdentifier(value string) *IdentifierNode {
	return &IdentifierNode{value: value}
}

func (n *IdentifierNode) Value() string {
	return n.value
}

func (n *IdentifierNode) IsQuoted() bool {
	return strings.HasPrefix(n.Value(), "`")
}

func (n *IdentifierNode) UnquotedValue() string {
	if n.IsQuoted() {
		return n.Value()
	}
	return n.value[1 : len(n.value)-1]
}

func (n *IdentifierNode) Children() []Node {
	return nil
}
