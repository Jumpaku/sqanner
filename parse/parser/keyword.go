package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseKeyword(s *ParseState) (node.KeywordNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	if !(s.ExpectNext(IsAnyKind(tokenize.TokenKeyword))) {
		return Error[node.KeywordNode](s, fmt.Errorf(`keyword not found`))
	}
	t := s.Next()

	return Accept(s, node.Keyword(node.KeywordCodeOf(string(t.Content))))
}
