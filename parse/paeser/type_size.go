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

	s.SkipSpaces()
	switch {
	default:
		return nil, parse.WrapError(s, parse.WrapError(s, fmt.Errorf(`invalid size: integer literal or 'MAX' not found`)))
	case s.ExpectNext(isKeyword("MAX")):
		return parse.Accept(s, node.TypeSizeMax()), nil
	case s.ExpectNext(isAnyKind(tokenize.TokenLiteralInteger)):
		t := s.Next()

		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return nil, parse.WrapError(s, parse.WrapError(s, fmt.Errorf(`invalid size: %w`, err)))
		}

		return parse.Accept(s, node.TypeSize(int(size))), nil
	}
}
