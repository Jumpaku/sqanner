package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseType(s *ParseState) (node.TypeNode, error) {
	Init(s)

	s.SkipSpacesAndComments()
	switch {
	default:
		return Error[node.TypeNode](s, fmt.Errorf(`invalid type: valid type identifier not found`))
	case s.ExpectNext(isIdentifier(true, `BOOL`)):
		s.Next()
		return Accept(s, node.BoolType())
	case s.ExpectNext(isIdentifier(true, `DATE`)):
		s.Next()
		return Accept(s, node.DateType())
	case s.ExpectNext(isIdentifier(true, `JSON`)):
		s.Next()
		return Accept(s, node.JSONType())
	case s.ExpectNext(isIdentifier(true, `INT64`)):
		s.Next()
		return Accept(s, node.Int64Type())
	case s.ExpectNext(isIdentifier(true, `NUMERIC`)):
		s.Next()
		return Accept(s, node.NumericType())
	case s.ExpectNext(isIdentifier(true, `FLOAT64`)):
		s.Next()
		return Accept(s, node.Float64Type())
	case s.ExpectNext(isIdentifier(true, `TIMESTAMP`)):
		s.Next()
		return Accept(s, node.TimestampType())
	case s.ExpectNext(isIdentifier(true, `BYTES`)):
		s.Next()
		s.SkipSpacesAndComments()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid type size: %w`, err))
		}
		return Accept(s, node.BytesType(size))
	case s.ExpectNext(isIdentifier(true, `STRING`)):
		s.Next()
		s.SkipSpacesAndComments()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return Error[node.TypeNode](s, fmt.Errorf(`invalid type size: %w`, err))
		}
		return Accept(s, node.StringType(size))
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
	}
}

func parseEnclosedTypeSize(s *ParseState) (node.TypeSizeNode, error) {
	s.SkipSpacesAndComments()
	if !s.ExpectNext(isSpecial('(')) {
		return Error[node.TypeSizeNode](s, fmt.Errorf(`invalid type size: '(' not found`))
	}
	s.Next()

	s.SkipSpacesAndComments()
	size, err := ParseTypeSize(s)
	if err != nil {
		return Error[node.TypeSizeNode](s, fmt.Errorf(`invalid type size: %w`, err))
	}

	s.SkipSpacesAndComments()
	if !s.ExpectNext(isSpecial(')')) {
		return Error[node.TypeSizeNode](s, fmt.Errorf(`invalid type size: ')' not found`))
	}
	s.Next()

	return size, nil
}
