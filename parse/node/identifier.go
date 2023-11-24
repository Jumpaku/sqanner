package node

import (
	"github.com/Jumpaku/sqanner/parse"
)

type IdentifierNode interface {
	Node
	isIdentifierNode()
	Value() string
	RequiresQuotes() bool
	QuotedValue() string
}

var _ IdentifierNode = (*identifierNode)(nil)

type identifierNode struct {
	NodeBase
	value          string
	requiresQuotes bool
}

func AcceptIdentifier(s *parse.ParseState, value string, requiresQuotes bool) *identifierNode {
	return &identifierNode{
		NodeBase:       NewNodeBase(s.Begin(), s.End()),
		value:          value,
		requiresQuotes: requiresQuotes,
	}
}

func NewIdentifier(value string, requiresQuotes bool) *identifierNode {
	return &identifierNode{value: value, requiresQuotes: requiresQuotes}
}

func (n *identifierNode) Value() string {
	return n.value
}

func (n *identifierNode) RequiresQuotes() bool {
	return n.requiresQuotes
}

func (n *identifierNode) QuotedValue() string {
	return "`" + n.value + "`"
}

func (n *identifierNode) Children() []Node {
	return nil
}

func (n *identifierNode) isIdentifierNode() {}
