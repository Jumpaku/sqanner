package types

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"github.com/Jumpaku/sqanner/tokenize"
	"strconv"
	"strings"
)

func ParseTypeSize(s *parse.ParseState) (types.TypeSizeNode, error) {
	s.Skip()
	t := s.PeekAt(0)
	switch {
	default:
		return nil, s.WrapError(fmt.Errorf(`integer literal or 'MAX' not found`))
	case parse.MatchIdentifier(t, true, `MAX`):
		s.Move(1)

		return types.AcceptTypeSizeMax(s), nil
	case parse.MatchAnyTokenKind(t, tokenize.TokenLiteralInteger):
		s.Move(1)

		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return nil, s.WrapError(fmt.Errorf(`fail to parse integer: %w`, err))
		}

		return types.AcceptTypeSize(s, int(size)), nil
	}
}
