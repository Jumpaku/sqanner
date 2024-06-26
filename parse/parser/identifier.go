package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *parse.ParseState) (node.IdentifierNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted) {
		return nil, s.WrapError(fmt.Errorf(`quoted or unquoted identifier not found`))
	}
	s.Move(1)

	unquoted := string(t.Content)
	requiresQuotes := t.Kind == tokenize.TokenIdentifierQuoted
	if requiresQuotes {
		unquoted = unquoted[1 : len(unquoted)-1]
	}

	return node.AcceptIdentifier(s, unquoted, requiresQuotes), nil
}
