package node

import "github.com/Jumpaku/sqanner/tokenize"

type nodeFunc[T Node] func(begin int, tokens []tokenize.Token) T

type Node interface {
	Kind() NodeKind
	Begin() int
	End() int
	Tokens() []tokenize.Token
	Children() []Node
}

type nodeBase struct {
	tokens []tokenize.Token
	begin  int
	kind   NodeKind
}

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
