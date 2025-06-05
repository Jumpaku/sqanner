package ddl

import (
	"fmt"
	"github.com/Jumpaku/sqanner/old_parse/node"
	"github.com/Jumpaku/sqanner/old_parse/parser"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseCreateTable(s *parser.ParseState) (node.IdentifierNode, error) {
	s.SkipSpacesAndComments()

	t := s.PeekAt(0)
	if !parser.IsAnyKind(t, tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted) {
		return nil, parser.Error(s, fmt.Errorf(`quoted or unquoted identifier not found`))
	}
	s.Move(1)

	return parser.Accept(s, node.Identifier(string(t.Content))), nil
}
