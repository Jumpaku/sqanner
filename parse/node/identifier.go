package node

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"strings"
)

type IdentifierNode interface {
	Node
	Value() string
	IsQuoted() bool
	UnquotedValue() string
}

func Identifier() NewNodeFunc[IdentifierNode] {
	return func(head int, tokens []tokenize.Token) IdentifierNode {
		return identifier{
			nodeBase{kind: NodeIdentifier, begin: head, tokens: tokens},
		}
	}
}

type identifier struct {
	nodeBase
}

var (
	_ Node           = identifier{}
	_ IdentifierNode = identifier{}
)

func (n identifier) Children() []Node {
	return nil
}

func (n identifier) Value() string {
	return string(n.Tokens()[0].Content)
}

func (n identifier) IsQuoted() bool {
	return strings.HasPrefix(n.Value(), "`")
}

func (n identifier) UnquotedValue() string {
	if n.IsQuoted() {
		return n.Value()
	}
	c := n.tokens[0].Content
	return string(c[1 : len(c)-1])
}
