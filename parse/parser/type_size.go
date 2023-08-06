package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
	"strconv"
	"strings"
)

func ParseTypeSize(s *ParseState) (node.TypeSizeNode, error) {
	s.SkipSpacesAndComments()
	t := s.PeekAt(0)
	switch {
	default:
		return nil, Error(s, fmt.Errorf(`integer literal or 'MAX' not found`))
	case IsIdentifier(t, true, `MAX`):
		s.Move(1)

		return Accept(s, node.TypeSizeMax()), nil
	case IsAnyKind(t, tokenize.TokenLiteralInteger):
		s.Move(1)

		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return nil, Error(s, fmt.Errorf(`fail to parse integer: %w`, err))
		}

		return Accept(s, node.TypeSize(int(size))), nil
	}
}
