package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/old_parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParsePath(s *ParseState) (node.PathNode, error) {
	s.SkipSpacesAndComments()

	var ch []node.IdentifierNode

	ident, err := ParseIdentifier(s.Child())
	if err != nil {
		return nil, Error(s, fmt.Errorf(`first identifier not found`))
	}
	s.Move(ident.Len())
	ch = append(ch, ident)

	for {
		s.SkipSpacesAndComments()
		switch {
		default:
			return Accept(s, node.Path(ch)), nil
		case IsAnySpecial(s.PeekAt(0), '.'):
			s.Move(1)

			s.SkipSpacesAndComments()
			switch {
			default:
				return nil, Error(s, fmt.Errorf(`identifier not found after '.'`))
			case IsAnyKind(s.PeekAt(0), tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted):
				n, err := ParseIdentifier(s.Child())
				if err != nil {
					return nil, Error(s, fmt.Errorf(`invalid identifier: %w`, err))
				}
				s.Move(n.Len())
				ch = append(ch, n)
			case IsAnyKind(s.PeekAt(0), tokenize.TokenKeyword):
				n, err := parseKeywordAsIdentifier(s.Child())
				if err != nil {
					return nil, Error(s, fmt.Errorf(`invalid keyword: %w`, err))
				}
				s.Move(n.Len())
				ch = append(ch, n)
			}
		}
	}
}

func parseKeywordAsIdentifier(s *ParseState) (node.IdentifierNode, error) {
	s.SkipSpacesAndComments()
	t := s.PeekAt(0)
	if !IsAnyKind(t, tokenize.TokenKeyword) {
		return nil, Error(s, fmt.Errorf(`keyword not found`))
	}
	s.Move(1)

	return Accept(s, node.Identifier(string(t.Content))), nil
}
