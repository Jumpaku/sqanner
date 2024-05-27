package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseKeyword(s *parse.ParseState) (node.KeywordNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenKeyword) {
		return nil, s.WrapError(fmt.Errorf(`keyword not found`))
	}
	s.Move(1)

	k, err := parse.OfKeyword(string(t.Content))
	if err != nil {
		return nil, s.WrapError(fmt.Errorf(`invalid keyword: %q`, string(t.Content)))
	}

	return node.AcceptKeyword(s, k), nil
}
