package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParsePath(s *parse.ParseState) (node.PathNode, error) {
	begin := s.Cursor
	s.Down()
	defer s.Up()
	s.SkipSpaces()

	var ch []node.IdentifierNode

	n, err := ParseIdentifier(s)
	if err != nil {
		return nil, err
	}
	ch = append(ch, n)

	isSeparator := func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenSpecialChar && t.Content[0] == '.'
	}

	for {
		switch {
		default:
			return node.Path(begin, s.Input[begin:s.Cursor], ch), nil
		case s.Expect(isSeparator):
			s.Next()

			n, err := ParseIdentifier(s)
			if err != nil {
				return nil, err
			}
			ch = append(ch, n)
		}
	}
}

func ParseIdentifier(s *parse.ParseState) (node.IdentifierNode, error) {
	begin := s.Cursor
	s.Down()
	defer s.Up()
	s.SkipSpaces()

	isIdentifier := func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenIdentifier || t.Kind == tokenize.TokenIdentifierQuoted
	}

	if !s.Expect(isIdentifier) {
		return nil, fmt.Errorf(`quoted or unquoted identifier is expected but not found`)
	}

	return node.Identifier(begin, s.Next()), nil
}
