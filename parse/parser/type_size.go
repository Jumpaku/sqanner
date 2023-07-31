package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
	"strconv"
	"strings"
)

func ParseTypeSize(s *ParseState) (node.TypeSizeNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	switch {
	default:
		return Error[node.TypeSizeNode](s, fmt.Errorf(`integer literal or 'MAX' not found`))
	case s.ExpectNext(isKeyword("MAX")):
		return Accept(s, node.TypeSizeMax())
	case s.ExpectNext(IsAnyKind(tokenize.TokenLiteralInteger)):
		t := s.Next()

		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return Error[node.TypeSizeNode](s, fmt.Errorf(`fail to parse integer: %w`, err))
		}

		return Accept(s, node.TypeSize(int(size)))
	}
}
