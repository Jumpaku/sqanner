package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *parse.ParseState) (node.IdentifierNode, error) {
	s.SkipSpacesAndComments()
	if !(s.ExpectNext(IsAnyKind(tokenize.TokenIdentifier)) || s.ExpectNext(IsAnyKind(tokenize.TokenIdentifierQuoted))) {
		return parse.WrapError[node.IdentifierNode](s, fmt.Errorf(`quoted or unquoted identifier not found`))
	}
	s.Next()

	return parse.Accept(s, node.Identifier())
}
