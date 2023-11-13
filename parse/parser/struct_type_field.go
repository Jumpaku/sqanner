package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseStructField(s *ParseState) (node.StructTypeFieldNode, error) {
	var err error
	var result node.StructTypeFieldNode
	switch {
	default:
		return nil, err
	case parseStructTypeFieldUnnamed(s.Child(), &result, &err):
		return result, nil
	case parseStructTypeFieldNamed(s.Child(), &result, &err):
		return result, nil
	}
}
func parseStructTypeFieldUnnamed(s *ParseState, out *node.StructTypeFieldNode, errOut *error) bool {
	s.SkipSpacesAndComments()
	typeNode, err := ParseType(s.Child())
	if err != nil {
		*errOut = Error(s, fmt.Errorf(`invalid type: %w`, err))
		return false
	}
	s.Move(typeNode.Len())
	*out = Accept(s, node.StructTypeFieldUnnamed(typeNode))
	return true
}

func parseStructTypeFieldNamed(s *ParseState, out *node.StructTypeFieldNode, errOut *error) bool {
	s.SkipSpacesAndComments()
	identNode, err := ParseIdentifier(s.Child())
	if err != nil {
		*errOut = Error(s, Error(s, fmt.Errorf(`invalid field name: %w`, err)))
		return false
	}
	s.Move(identNode.Len())

	s.SkipSpacesAndComments()
	var typeNode node.TypeNode
	typeNode, err = ParseType(s.Child())
	if err != nil {
		*errOut = Error(s, Error(s, fmt.Errorf(`invalid field type: %w`, err)))
		return false
	}
	s.Move(typeNode.Len())
	*out = Accept(s, node.StructTypeField(identNode, typeNode))
	return true
}
