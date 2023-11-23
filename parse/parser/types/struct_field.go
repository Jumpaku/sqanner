package types

import (
	"errors"
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"github.com/Jumpaku/sqanner/parse/parser"
)

func ParseStructField(s *parse.ParseState) (types.StructFieldNode, error) {
	s.Skip()

	var err error
	if n, e := parseStructFieldUnnamed(s.Child()); e != nil {
		err = errors.Join(err, fmt.Errorf("invalid struct field: %w", e))
	} else {
		s.Move(n.Len())
		return n, nil
	}

	if n, e := parseStructFieldNamed(s.Child()); e != nil {
		err = errors.Join(err, fmt.Errorf("invalid struct field: %w", e))
	} else {
		s.Move(n.Len())
		return n, nil
	}

	return nil, s.WrapError(err)
}
func parseStructFieldUnnamed(s *parse.ParseState) (types.StructFieldNode, error) {
	s.Skip()

	typeNode, err := ParseType(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf(`invalid type: %w`, err))
	}
	s.Move(typeNode.Len())
	return types.AcceptStructFieldUnnamed(s, typeNode), nil
}

func parseStructFieldNamed(s *parse.ParseState) (types.StructFieldNode, error) {
	s.Skip()

	identNode, err := parser.ParseIdentifier(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf(`invalid field name: %w`, err))
	}
	s.Move(identNode.Len())

	s.Skip()
	var typeNode types.TypeNode
	typeNode, err = ParseType(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf(`invalid field type: %w`, err))
	}
	s.Move(typeNode.Len())
	return types.AcceptStructField(s, identNode, typeNode), nil
}
