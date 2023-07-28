package paeser

import (
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
	"strconv"
	"strings"
)

func ParseType(s *parse.ParseState) (node.TypeNode, error) {
	begin := s.Cursor
	s.Down()
	defer s.Up()

	s.SkipSpaces()
	switch {
	default:
		return nil, fmt.Errorf(`invalid type`)
	case s.Expect(isIdentifier(`BOOL`)):
		return node.BoolType(begin, s.Next()), nil
	case s.Expect(isIdentifier(`DATE`)):
		return node.DateType(begin, s.Next()), nil
	case s.Expect(isIdentifier(`JSON`)):
		return node.JSONType(begin, s.Next()), nil
	case s.Expect(isIdentifier(`INT64`)):
		return node.Int64Type(begin, s.Next()), nil
	case s.Expect(isIdentifier(`NUMERIC`)):
		return node.NumericType(begin, s.Next()), nil
	case s.Expect(isIdentifier(`FLOAT64`)):
		return node.Float64Type(begin, s.Next()), nil
	case s.Expect(isIdentifier(`TIMESTAMP`)):
		return node.TimestampType(begin, s.Next()), nil
	case s.Expect(isKeyword(`Array`)):
		s.Next()

		s.SkipSpaces()
		if !s.Expect(isSpecial('<')) {
			return nil, fmt.Errorf(`invalid type parameter: '<' not found`)
		}
		s.Next()

		s.SkipSpaces()
		element, err := ParseType(s)
		if err != nil {
			return nil, fmt.Errorf(`invalid type parameter: %w`, err)
		}

		s.SkipSpaces()
		if !s.Expect(isSpecial('>')) {
			return nil, fmt.Errorf(`invalid type parameter: '>' not found`)
		}
		s.Next()

		return node.ArrayType(begin, s.Input[begin:s.Cursor], element), nil

	case s.Expect(isKeyword(`STRUCT`)):
		s.Next()

		s.SkipSpaces()
		if !s.Expect(isSpecial('<')) {
			return nil, fmt.Errorf(`invalid type parameter: '<' not found`)
		}
		s.Next()

		var fields []node.StructTypeFieldNode
		for {
			s.SkipSpaces()
			field, err := ParseStructField(s)
			if err != nil {
				return nil, fmt.Errorf(`invalid struct field: %w`, err)
			}
			fields = append(fields, field)

			s.SkipSpaces()
			switch {
			default:
				return nil, fmt.Errorf(`invalid type parameter: ',' or '>' not found`)
			case s.Expect(isSpecial(',')):
				s.Next()
			case s.Expect(isSpecial('>')):
				s.Next()
				return node.StructType(begin, s.Input[begin:s.Cursor], fields), nil
			}
		}

	case s.Expect(isIdentifier(`BYTES`)), s.Expect(isIdentifier(`STRING`)):
		t := s.Next()

		s.SkipSpaces()
		if !s.Expect(isSpecial('(')) {
			return nil, fmt.Errorf(`invalid type parameter: '(' not found`)
		}
		s.Next()

		s.SkipSpaces()
		size, err := ParseTypeSize(s)
		if err != nil {
			return nil, fmt.Errorf(`invalid size: %w`, err)
		}

		s.SkipSpaces()
		if !s.Expect(isSpecial('>')) {
			return nil, fmt.Errorf(`invalid type parameter: '>' not found`)
		}
		s.Next()

		switch {
		default:
			return assert.Unexpected2[node.TypeNode, error](``)
		case isIdentifier(`BYTES`)(t):
			return node.BytesType(begin, s.Input[begin:s.Cursor], size), nil
		case isIdentifier(`STRING`)(t):
			return node.StringType(begin, s.Input[begin:s.Cursor], size), nil
		}
	}
}

func ParseStructField(s *parse.ParseState) (node.StructTypeFieldNode, error) {
	begin := s.Cursor
	s.Down()
	defer s.Up()

	var fieldName node.IdentifierNode

	s.SkipSpaces()
	if s.Expect(isKind(tokenize.TokenIdentifier)) || s.Expect(isKind(tokenize.TokenIdentifierQuoted)) {
		var err error
		fieldName, err = ParseIdentifier(s)
		if err != nil {
			return nil, fmt.Errorf(`invalid field name: %w`, err)
		}
	}

	s.SkipSpaces()
	fieldType, err := ParseType(s)
	if err != nil {
		return nil, fmt.Errorf(`invalid field type: %w`, err)
	}

	return node.StructTypeField(begin, s.Input[begin:s.Cursor], fieldName, fieldType), nil
}

func ParseTypeSize(s *parse.ParseState) (node.TypeSizeNode, error) {
	begin := s.Cursor
	s.Down()
	defer s.Up()

	s.SkipSpaces()
	switch {
	default:
		return nil, fmt.Errorf(`invalid type parameter: integer literal or 'MAX' not found`)
	case s.Expect(isKeyword("MAX")):
		return node.TypeSizeMax(begin, s.Next()), nil
	case s.Expect(isKind(tokenize.TokenLiteralInteger)):
		t := s.Next()
		size, err := strconv.ParseInt(strings.ToLower(string(t.Content)), 0, 64)
		if err != nil {
			return nil, fmt.Errorf(`invalid type size: %w`, err)
		}

		return node.TypeSize(begin, t, int(size)), nil
	}
}
