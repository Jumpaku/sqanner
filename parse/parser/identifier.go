package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *parse.ParseState) (*node.IdentifierNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.IsAnyKind(t, tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted) {
		return nil, s.WrapError(fmt.Errorf(`quoted or unquoted identifier not found`))
	}

	content := string(t.Content)
	if parse.IsAKeyword(content) {
		return nil, s.WrapError(fmt.Errorf(`keyword is not identifier`))
	}
	s.Move(1)

	return node.AcceptIdentifier(s, content), nil
}
