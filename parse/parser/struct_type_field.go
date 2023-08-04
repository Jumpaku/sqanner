package parser

import (
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseStructField(s *ParseState) (node.StructTypeFieldNode, error) {
	Init(s)
	/*
		s.SkipSpacesAndComments()
		if !s.ExpectNext(IsAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted)) {
			return Error[node.StructTypeFieldNode](s, fmt.Errorf(`identifier not found`))
		}
		firstIdent, err := ParseIdentifier(s)

		s.SkipSpacesAndComments()
		hasFieldName := s.ExpectNext(isSpecial(','))IsAnyKind(tokenize.TokenIdentifier, tokenize.TokenIdentifierQuoted))
		var secondIdent node.IdentifierNode
		if hasFieldName {
			secondIdent, err = ParseIdentifier(s)
			if err != nil {
				return Error[node.StructTypeFieldNode](s, fmt.Errorf(`invalid field name: %w`, err))
			}
		}

		if hasFieldName {
			return Expect(s, node.StructTypeField(firstIdent, secondIdent))
		}
		var fieldName node.IdentifierNode
		if !anonymous {
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

		if anonymous {
			return Expect(s, node.StructTypeFieldAnonymous(fieldType))
		}

		return Expect(s, node.StructTypeField(fieldName, fieldType))
	*/
	return nil, nil
}
