package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *ParseState) (node.IdentifierNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	if !(s.ExpectNext(IsAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted))) {
		return Error[node.IdentifierNode](s, fmt.Errorf(`quoted or unquoted identifier not found`))
	}
	t := s.Next()

	return Accept(s, node.Identifier(string(t.Content)))
}
