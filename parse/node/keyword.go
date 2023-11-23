package node

import (
	"github.com/Jumpaku/sqanner/parse"
)

type KeywordNode interface {
	Node
	KeywordCode() parse.KeywordCode
}

var _ KeywordNode = (*keywordNode)(nil)

type keywordNode struct {
	NodeBase
	code parse.KeywordCode
}

func AcceptKeyword(s *parse.ParseState, code parse.KeywordCode) *keywordNode {
	return &keywordNode{
		NodeBase: NewNodeBase(s.Begin(), s.End()),
		code:     code,
	}
}

func NewKeyword(code parse.KeywordCode) *keywordNode {
	return &keywordNode{code: code}
}

func (n *keywordNode) KeywordCode() parse.KeywordCode {
	return n.code
}

func (n *keywordNode) Children() []Node {
	return nil
}
