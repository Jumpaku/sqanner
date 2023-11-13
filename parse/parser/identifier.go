package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *ParseState) (node.IdentifierNode, error) {
	s.SkipSpacesAndComments()

	t := s.PeekAt(0)
	if !IsAnyKind(t, tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted) {
		return nil, Error(s, fmt.Errorf(`quoted or unquoted identifier not found`))
	}
	s.Move(1)

	return Accept(s, node.Identifier(string(t.Content))), nil
}
