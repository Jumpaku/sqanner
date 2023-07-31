package paeser

import (
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseType(s *parse.ParseState) (node.TypeNode, error) {
	parse.Stack(s)

	s.SkipSpacesAndComments()
	switch {
	default:
		return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid type: valid type identifier not found`))
	case s.ExpectNext(isIdentifier(`BOOL`)):
		return parse.Accept(s, node.BoolType())
	case s.ExpectNext(isIdentifier(`DATE`)):
		return parse.Accept(s, node.DateType())
	case s.ExpectNext(isIdentifier(`JSON`)):
		return parse.Accept(s, node.JSONType())
	case s.ExpectNext(isIdentifier(`INT64`)):
		return parse.Accept(s, node.Int64Type())
	case s.ExpectNext(isIdentifier(`NUMERIC`)):
		return parse.Accept(s, node.NumericType())
	case s.ExpectNext(isIdentifier(`FLOAT64`)):
		return parse.Accept(s, node.Float64Type())
	case s.ExpectNext(isIdentifier(`TIMESTAMP`)):
		return parse.Accept(s, node.TimestampType())
	case s.ExpectNext(isKeyword(`Array`)):
		s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('<')) {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid array type element: '<' not found`))
		}
		s.Next()

		s.SkipSpacesAndComments()
		element, err := ParseType(s)
		if err != nil {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid array type element: %w`, err))
		}

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('>')) {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid array type element: '>' not found`))
		}
		s.Next()

		return parse.Accept(s, node.ArrayType(element))

	case s.ExpectNext(isKeyword(`STRUCT`)):
		s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('<')) {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid struct type field: '<' not found`))
		}
		s.Next()

		var fields []node.StructTypeFieldNode
		for {
			s.SkipSpacesAndComments()
			field, err := ParseStructField(s)
			if err != nil {
				return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid struct type field: %w`, err))
			}
			fields = append(fields, field)

			s.SkipSpacesAndComments()
			switch {
			default:
				return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid struct type field: ',' or '>' not found`))
			case s.ExpectNext(isSpecial(',')):
				s.Next()
			case s.ExpectNext(isSpecial('>')):
				s.Next()
				return parse.Accept(s, node.StructType(fields))
			}
		}

	case s.ExpectNext(isIdentifier(`BYTES`)), s.ExpectNext(isIdentifier(`STRING`)):
		t := s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('(')) {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid type size: '(' not found`))
		}
		s.Next()

		s.SkipSpacesAndComments()
		size, err := ParseTypeSize(s)
		if err != nil {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid type size: %w`, err))
		}

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial(')')) {
			return parse.WrapError[node.TypeNode](s, fmt.Errorf(`invalid type size: ')' not found`))
		}
		s.Next()

		switch {
		default:
			return assert.Unexpected2[node.TypeNode, error](``)
		case isIdentifier(`BYTES`)(t):
			return parse.Accept(s, node.BytesType(size))
		case isIdentifier(`STRING`)(t):
			return parse.Accept(s, node.StringType(size))
		}
	}
}
