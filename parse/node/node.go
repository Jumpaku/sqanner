package node

import "github.com/Jumpaku/sqanner/tokenize"

type Node interface {
	Kind() NodeKind
	Begin() int
	End() int
	Tokens() []tokenize.Token
}

type TerminalNode interface {
	Node
	Token() tokenize.Token
}

type nodeBase struct {
	tokens []tokenize.Token
	begin  int
	kind   NodeKind
}

var _ Node = nodeBase{}

func (n nodeBase) Kind() NodeKind {
	return n.kind
}

func (n nodeBase) Begin() int {
	return n.begin
}

func (n nodeBase) End() int {
	return n.begin + len(n.tokens)
}

func (n nodeBase) Tokens() []tokenize.Token {
	return n.tokens
}
