package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParsePath(s *ParseState) (node.PathNode, error) {
	Init(s)

	var ch []node.IdentifierNode

	s.SkipSpacesAndComments()
	n, err := ParseIdentifier(s)
	if err != nil {
		return Error[node.PathNode](s, fmt.Errorf(`first identifier not found`))
	}
	ch = append(ch, n)

	isSeparator := func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenSpecialChar && t.Content[0] == '.'
	}

	for {
		s.SkipSpacesAndComments()
		switch {
		default:
			return Expect(s, node.Path(ch))
		case s.ExpectNext(isSeparator):
			s.Next()

			s.SkipSpacesAndComments()
			switch {
			default:
				return Error[node.PathNode](s, fmt.Errorf(`identifier not found after '.'`))
			case s.ExpectNext(IsAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted)):
				n, err := ParseIdentifier(s)
				if err != nil {
					return Error[node.PathNode](s, fmt.Errorf(`invalid identifier: %w`, err))
				}
				ch = append(ch, n)
			case s.ExpectNext(IsAnyKind(tokenize.TokenKeyword)):
				n, err := parseKeywordAsIdentifier(s)
				if err != nil {
					return Error[node.PathNode](s, fmt.Errorf(`invalid keyword: %w`, err))
				}
				ch = append(ch, n)
			}
		}
	}
}

func parseKeywordAsIdentifier(s *ParseState) (node.IdentifierNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	if !(s.ExpectNext(IsAnyKind(tokenize.TokenKeyword))) {
		return Error[node.IdentifierNode](s, fmt.Errorf(`keyword not found`))
	}
	t := s.Next()

	return Expect(s, node.Identifier(string(t.Content)))
}
