package paeser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseStructField(s *parse.ParseState) (node.StructTypeFieldNode, error) {
	parse.Stack(s)

	var fieldName node.IdentifierNode

	s.SkipSpaces()
	if s.ExpectNext(isAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted)) {
		var err error
		fieldName, err = ParseIdentifier(s)
		if err != nil {
			return nil, parse.WrapError(s, parse.WrapError(s, fmt.Errorf(`invalid struct field: invalid field name: %w`, err)))
		}
	}

	s.SkipSpaces()
	fieldType, err := ParseType(s)
	if err != nil {
		return nil, parse.WrapError(s, parse.WrapError(s, fmt.Errorf(`invalid struct field: invalid field type: %w`, err)))
	}

	return parse.Accept(s, node.StructTypeField(fieldName, fieldType)), nil
}
