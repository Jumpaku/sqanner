package node

import (
	"strings"
)

type IdentifierNode interface {
	Node
	Value() string
	IsQuoted() bool
	UnquotedValue() string
}

func Identifier(value string) NewNodeFunc[IdentifierNode] {
	return func(begin, end int) IdentifierNode {
		return identifier{
			nodeBase: nodeBase{kind: KindIdentifier, begin: begin, end: end},
			value:    value,
		}
	}
}

type identifier struct {
	nodeBase
	value string
}

var (
	_ Node           = identifier{}
	_ IdentifierNode = identifier{}
)

func (n identifier) Children() []Node {
	return nil
}

func (n identifier) Value() string {
	return n.value
}

func (n identifier) IsQuoted() bool {
	return strings.HasPrefix(n.Value(), "`")
}

func (n identifier) UnquotedValue() string {
	if n.IsQuoted() {
		return n.Value()
	}
	return n.value[1 : len(n.value)-1]
}
