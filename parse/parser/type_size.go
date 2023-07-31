package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
	"strconv"
	"strings"
)

func ParseTypeSize(s *parse.ParseState) (node.TypeSizeNode, error) {
	parse.Stack(s)

	s.SkipSpacesAndComments()
	switch {
	default:
		return parse.WrapError[node.TypeSizeNode](s, fmt.Errorf(`integer literal or 'MAX' not found`))
	case s.ExpectNext(isKeyword("MAX")):
		return parse.Accept(s, node.TypeSizeMax())
	case s.ExpectNext(IsAnyKind(tokenize.TokenLiteralInteger)):
		t := s.Next()

		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return parse.WrapError[node.TypeSizeNode](s, fmt.Errorf(`fail to parse integer: %w`, err))
		}

		return parse.Accept(s, node.TypeSize(int(size)))
	}
}
