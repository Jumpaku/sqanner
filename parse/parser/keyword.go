package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseKeyword(s *ParseState) (node.KeywordNode, error) {
	s.SkipSpacesAndComments()

	t := s.PeekAt(0)
	if !IsAnyKind(t, tokenize.TokenKeyword) {
		return nil, Error(s, fmt.Errorf(`keyword not found`))
	}

	return Accept(s, node.Keyword(node.KeywordCodeOf(string(t.Content)))), nil
}
