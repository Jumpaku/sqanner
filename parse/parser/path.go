package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParsePath(s *parse.ParseState) (node.PathNode, error) {
	var ch []node.IdentifierNode

	s.SkipSpacesAndComments()
	n, err := ParseIdentifier(s)
	if err != nil {
		return parse.WrapError[node.PathNode](s, fmt.Errorf(`first identifier not found`))
	}
	ch = append(ch, n)

	isSeparator := func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenSpecialChar && t.Content[0] == '.'
	}

	for {
		switch {
		default:
			return parse.Accept(s, node.Path(ch))
		case s.ExpectNext(isSeparator):
			s.Next()

			n, err := ParseIdentifier(s)
			if err != nil {
				return parse.WrapError[node.PathNode](s, fmt.Errorf(`identifier not found after '.'`))
			}
			ch = append(ch, n)
		}
	}
}
