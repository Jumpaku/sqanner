package node

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"strings"
)

type IdentifierNode interface {
	TerminalNode
	Value() string
	IsQuoted() bool
	UnquotedValue() string
}

type PathNode interface {
	Node
	Identifiers() []IdentifierNode
}

func Identifier(begin int, token tokenize.Token) IdentifierNode {
	return identifier{
		nodeBase{kind: NodeIdentifier, begin: begin, tokens: []tokenize.Token{token}},
	}
}

func Path(begin int, tokens []tokenize.Token, identifiers []IdentifierNode) PathNode {
	return path{
		nodeBase:    nodeBase{kind: NodePath, begin: begin, tokens: tokens},
		identifiers: identifiers,
	}
}

type identifier struct {
	nodeBase
}

var (
	_ Node           = identifier{}
	_ TerminalNode   = identifier{}
	_ IdentifierNode = identifier{}
)

func (n identifier) Token() tokenize.Token {
	return n.Tokens()[0]
}

func (n identifier) Value() string {
	return string(n.Token().Content)
}

func (n identifier) IsQuoted() bool {
	return strings.HasPrefix(n.Value(), "`")
}

func (n identifier) UnquotedValue() string {
	if n.IsQuoted() {
		return n.Value()
	}
	return strings.TrimSuffix(strings.TrimPrefix(n.Value(), "`"), "`")
}

type path struct {
	nodeBase
	identifiers []IdentifierNode
}

var _ Node = path{}
var _ PathNode = path{}

func (n path) Identifiers() []IdentifierNode {
	return n.identifiers
}
