package parser

import (
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseType(s *ParseState) (node.TypeNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	switch {
	default:
		return Error[node.TypeNode](s, fmt.Errorf(`invalid type: valid type identifier not found`))
	case s.ExpectNext(isIdentifier(`BOOL`)):
		return Accept(s, node.BoolType())
	case s.ExpectNext(isIdentifier(`DATE`)):
		return Accept(s, node.DateType())
	case s.ExpectNext(isIdentifier(`JSON`)):
		return Accept(s, node.JSONType())
	case s.ExpectNext(isIdentifier(`INT64`)):
		return Accept(s, node.Int64Type())
	case s.ExpectNext(isIdentifier(`NUMERIC`)):
		return Accept(s, node.NumericType())
	case s.ExpectNext(isIdentifier(`FLOAT64`)):
		return Accept(s, node.Float64Type())
	case s.ExpectNext(isIdentifier(`TIMESTAMP`)):
		return Accept(s, node.TimestampType())
	case s.ExpectNext(isKeyword(`Array`)):
		s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('<')) {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid array type element: '<' not found`))
		}
		s.Next()

		s.SkipSpacesAndComments()
		element, err := ParseType(s)
		if err != nil {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid array type element: %w`, err))
		}

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('>')) {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid array type element: '>' not found`))
		}
		s.Next()

		return Accept(s, node.ArrayType(element))

	case s.ExpectNext(isKeyword(`STRUCT`)):
		s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('<')) {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid struct type field: '<' not found`))
		}
		s.Next()

		var fields []node.StructTypeFieldNode
		for {
			s.SkipSpacesAndComments()
			field, err := ParseStructField(s)
			if err != nil {
				return Error[node.TypeNode](s, fmt.Errorf(`invalid struct type field: %w`, err))
			}
			fields = append(fields, field)

			s.SkipSpacesAndComments()
			switch {
			default:
				return Error[node.TypeNode](s, fmt.Errorf(`invalid struct type field: ',' or '>' not found`))
			case s.ExpectNext(isSpecial(',')):
				s.Next()
			case s.ExpectNext(isSpecial('>')):
				s.Next()
				return Accept(s, node.StructType(fields))
			}
		}

	case s.ExpectNext(isIdentifier(`BYTES`)), s.ExpectNext(isIdentifier(`STRING`)):
		t := s.Next()

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial('(')) {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid type size: '(' not found`))
		}
		s.Next()

		s.SkipSpacesAndComments()
		size, err := ParseTypeSize(s)
		if err != nil {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid type size: %w`, err))
		}

		s.SkipSpacesAndComments()
		if !s.ExpectNext(isSpecial(')')) {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid type size: ')' not found`))
		}
		s.Next()

		switch {
		default:
			return assert.Unexpected2[node.TypeNode, error](``)
		case isIdentifier(`BYTES`)(t):
			return Accept(s, node.BytesType(size))
		case isIdentifier(`STRING`)(t):
			return Accept(s, node.StringType(size))
		}
	}
}
