package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/old_parse/node"
)

func ParseType(s *ParseState) (node.TypeNode, error) {
	s.SkipSpacesAndComments()
	switch {
	default:
		return nil, Error(s, fmt.Errorf(`invalid type: valid type identifier not found`))
	case IsIdentifier(s.PeekAt(0), true, `BOOL`):
		s.Move(1)
		return Accept(s, node.BoolType()), nil
	case IsIdentifier(s.PeekAt(0), true, `DATE`):
		s.Move(1)
		return Accept(s, node.DateType()), nil
	case IsIdentifier(s.PeekAt(0), true, `JSON`):
		s.Move(1)
		return Accept(s, node.JSONType()), nil
	case IsIdentifier(s.PeekAt(0), true, `INT64`):
		s.Move(1)
		return Accept(s, node.Int64Type()), nil
	case IsIdentifier(s.PeekAt(0), true, `NUMERIC`):
		s.Move(1)
		return Accept(s, node.NumericType()), nil
	case IsIdentifier(s.PeekAt(0), true, `FLOAT64`):
		s.Move(1)
		return Accept(s, node.Float64Type()), nil
	case IsIdentifier(s.PeekAt(0), true, `TIMESTAMP`):
		s.Move(1)
		return Accept(s, node.TimestampType()), nil
	case IsIdentifier(s.PeekAt(0), true, `BYTES`):
		s.Move(1)
		s.SkipSpacesAndComments()
		if !IsAnySpecial(s.PeekAt(0), '(') {
			return Accept(s, node.BytesType()), nil
		}

		s.SkipSpacesAndComments()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return nil, Error(s, fmt.Errorf(`invalid type size: %w`, err))
		}
		return Accept(s, node.BytesTypeSized(size)), nil
	case IsIdentifier(s.PeekAt(0), true, `STRING`):
		s.Move(1)
		s.SkipSpacesAndComments()
		if !IsAnySpecial(s.PeekAt(0), '(') {
			return Accept(s, node.StringType()), nil
		}

		s.SkipSpacesAndComments()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return nil, Error(s, fmt.Errorf(`invalid type size: %w`, err))
		}
		return Accept(s, node.StringTypeSized(size)), nil
	case IsKeyword(s.PeekAt(0), node.KeywordArray):
		s.Move(1)

		s.SkipSpacesAndComments()
		if !IsAnySpecial(s.PeekAt(0), '<') {
			return nil, Error(s, fmt.Errorf(`invalid array type element: '<' not found`))
		}
		s.Move(1)

		s.SkipSpacesAndComments()
		element, err := ParseType(s.Child())
		if err != nil {
			return nil, Error(s, fmt.Errorf(`invalid array type element: %w`, err))
		}
		s.Move(element.Len())

		s.SkipSpacesAndComments()
		if !IsAnySpecial(s.PeekAt(0), '>') {
			return nil, Error(s, fmt.Errorf(`invalid array type element: '>' not found`))
		}
		s.Move(1)

		return Accept(s, node.ArrayType(element)), nil

	case IsKeyword(s.PeekAt(0), node.KeywordStruct):
		s.Move(1)

		s.SkipSpacesAndComments()
		if !IsAnySpecial(s.PeekAt(0), '<') {
			return nil, Error(s, fmt.Errorf(`invalid struct type field: '<' not found`))
		}
		s.Move(1)

		s.SkipSpacesAndComments()
		if IsAnySpecial(s.PeekAt(0), '>') {
			s.Move(1)
			return Accept(s, node.StructType(nil)), nil
		}

		var fields []node.StructTypeFieldNode
		for {
			field, err := ParseStructField(s.Child())
			if err != nil {
				return nil, Error(s, fmt.Errorf(`invalid struct type field: %w`, err))
			}
			s.Move(field.Len())
			fields = append(fields, field)

			s.SkipSpacesAndComments()
			switch {
			default:
				return nil, Error(s, fmt.Errorf(`invalid struct type field: ',' or '>' not found`))
			case IsAnySpecial(s.PeekAt(0), ','):
				s.Move(1)
			case IsAnySpecial(s.PeekAt(0), '>'):
				s.Move(1)
				return Accept(s, node.StructType(fields)), nil
			}
		}
	}
}

func parseEnclosedTypeSize(s *ParseState) (node.TypeSizeNode, error) {
	s.SkipSpacesAndComments()
	if !IsAnySpecial(s.PeekAt(0), '(') {
		return nil, Error(s, fmt.Errorf(`invalid type size: '(' not found`))
	}
	s.Move(1)

	s.SkipSpacesAndComments()
	size, err := ParseTypeSize(s.Child())
	if err != nil {
		return nil, Error(s, fmt.Errorf(`invalid type size: %w`, err))
	}
	s.Move(size.Len())

	s.SkipSpacesAndComments()
	if !IsAnySpecial(s.PeekAt(0), ')') {
		return nil, Error(s, fmt.Errorf(`invalid type size: ')' not found`))
	}
	s.Move(1)

	return size, nil
}
