package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
)

func ParseStructField(s *ParseState) (node.StructTypeFieldNode, error) {
	Init(s)

	var fieldName node.IdentifierNode

	s.SkipSpacesAndComments()
	if s.ExpectNext(IsAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted)) {
		var err error
		fieldName, err = ParseIdentifier(s)
		if err != nil {
			return Error[node.StructTypeFieldNode](s, fmt.Errorf(`invalid field name: %w`, err))
		}
	}

	s.SkipSpacesAndComments()
	fieldType, err := ParseType(s)
	if err != nil {
		return Error[node.StructTypeFieldNode](s, fmt.Errorf(`invalid field type: %w`, err))
	}

	return Accept(s, node.StructTypeField(fieldName, fieldType))
}
