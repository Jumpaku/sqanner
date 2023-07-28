package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseIdentifier(s *parse.ParseState) (node.IdentifierNode, error) {
	s.SkipSpaces()
	if !(s.ExpectNext(isAnyKind(tokenize.TokenIdentifier)) || s.ExpectNext(isAnyKind(tokenize.TokenIdentifierQuoted))) {
		return nil, parse.WrapError(s, fmt.Errorf(`invalid identifier: quoted or unquoted identifier not found`))
	}
	s.Next()

	return parse.Accept(s, node.Identifier()), nil
}
