package paeser

import (
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
)

func ParseType(s *parse.ParseState) (node.TypeNode, error) {
	parse.Stack(s)

	s.SkipSpaces()
	switch {
	default:
		return nil, parse.WrapError(s, fmt.Errorf(`invalid type: valid type identifier not found`))
	case s.ExpectNext(isIdentifier(`BOOL`)):
		return parse.Accept(s, node.BoolType()), nil
	case s.ExpectNext(isIdentifier(`DATE`)):
		return parse.Accept(s, node.DateType()), nil
	case s.ExpectNext(isIdentifier(`JSON`)):
		return parse.Accept(s, node.JSONType()), nil
	case s.ExpectNext(isIdentifier(`INT64`)):
		return parse.Accept(s, node.Int64Type()), nil
	case s.ExpectNext(isIdentifier(`NUMERIC`)):
		return parse.Accept(s, node.NumericType()), nil
	case s.ExpectNext(isIdentifier(`FLOAT64`)):
		return parse.Accept(s, node.Float64Type()), nil
	case s.ExpectNext(isIdentifier(`TIMESTAMP`)):
		return parse.Accept(s, node.TimestampType()), nil
	case s.ExpectNext(isKeyword(`Array`)):
		s.Next()

		s.SkipSpaces()
		if !s.ExpectNext(isSpecial('<')) {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid array element type: '<' not found`))
		}
		s.Next()

		s.SkipSpaces()
		element, err := ParseType(s)
		if err != nil {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid array element type: %w`, err))
		}

		s.SkipSpaces()
		if !s.ExpectNext(isSpecial('>')) {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid array element type: '>' not found`))
		}
		s.Next()

		return parse.Accept(s, node.ArrayType(element)), nil

	case s.ExpectNext(isKeyword(`STRUCT`)):
		s.Next()

		s.SkipSpaces()
		if !s.ExpectNext(isSpecial('<')) {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid struct field: '<' not found`))
		}
		s.Next()

		var fields []node.StructTypeFieldNode
		for {
			s.SkipSpaces()
			field, err := ParseStructField(s)
			if err != nil {
				return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid struct field: %w`, err))
			}
			fields = append(fields, field)

			s.SkipSpaces()
			switch {
			default:
				return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid struct field: ',' or '>' not found`))
			case s.ExpectNext(isSpecial(',')):
				s.Next()
			case s.ExpectNext(isSpecial('>')):
				s.Next()
				return parse.Accept(s, node.StructType(fields)), nil
			}
		}

	case s.ExpectNext(isIdentifier(`BYTES`)), s.ExpectNext(isIdentifier(`STRING`)):
		t := s.Next()

		s.SkipSpaces()
		if !s.ExpectNext(isSpecial('(')) {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid size: '(' not found`))
		}
		s.Next()

		s.SkipSpaces()
		size, err := ParseTypeSize(s)
		if err != nil {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid size: %w`, err))
		}

		s.SkipSpaces()
		if !s.ExpectNext(isSpecial(')')) {
			return nil, parse.WrapError(s, fmt.Errorf(`invalid type: invalid size: ')' not found`))
		}
		s.Next()

		switch {
		default:
			return assert.Unexpected2[node.TypeNode, error](``)
		case isIdentifier(`BYTES`)(t):
			return parse.Accept(s, node.BytesType(size)), nil
		case isIdentifier(`STRING`)(t):
			return parse.Accept(s, node.StringType(size)), nil
		}
	}
}
