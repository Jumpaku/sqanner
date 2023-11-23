package types

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node/types"
)

func ParseType(s *parse.ParseState) (types.TypeNode, error) {
	s.Skip()
	switch {
	default:
		return nil, s.WrapError(fmt.Errorf(`invalid type: valid type identifier not found`))
	case parse.MatchIdentifier(s.PeekAt(0), true, `BOOL`):
		s.Move(1)
		return types.AcceptBool(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `DATE`):
		s.Move(1)
		return types.AcceptDate(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `JSON`):
		s.Move(1)
		return types.AcceptJSON(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `INT64`):
		s.Move(1)
		return types.AcceptInt64(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `NUMERIC`):
		s.Move(1)
		return types.AcceptNumeric(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `FLOAT64`):
		s.Move(1)
		return types.AcceptFloat64(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `TIMESTAMP`):
		s.Move(1)
		return types.AcceptTimestamp(s), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `BYTES`):
		s.Move(1)
		s.Skip()
		if !parse.MatchAnySpecial(s.PeekAt(0), '(') {
			return types.AcceptBytes(s), nil
		}

		s.Skip()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return nil, s.WrapError(fmt.Errorf(`invalid type size: %w`, err))
		}
		return types.AcceptBytesSized(s, size), nil
	case parse.MatchIdentifier(s.PeekAt(0), true, `STRING`):
		s.Move(1)
		s.Skip()
		if !parse.MatchAnySpecial(s.PeekAt(0), '(') {
			return types.AcceptString(s), nil
		}

		s.Skip()
		size, err := parseEnclosedTypeSize(s)
		if err != nil {
			return nil, s.WrapError(fmt.Errorf(`invalid type size: %w`, err))
		}
		return types.AcceptStringSized(s, size), nil
	case parse.MatchAnyKeyword(s.PeekAt(0), parse.KeywordCodeArray):
		s.Move(1)

		s.Skip()
		if !parse.MatchAnySpecial(s.PeekAt(0), '<') {
			return nil, s.WrapError(fmt.Errorf(`invalid array type element: '<' not found`))
		}
		s.Move(1)

		s.Skip()
		element, err := ParseType(s.Child())
		if err != nil {
			return nil, s.WrapError(fmt.Errorf(`invalid array type element: %w`, err))
		}
		s.Move(element.Len())

		s.Skip()
		if !parse.MatchAnySpecial(s.PeekAt(0), '>') {
			return nil, s.WrapError(fmt.Errorf(`invalid array type element: '>' not found`))
		}
		s.Move(1)

		return types.AcceptArray(s, element), nil

	case parse.MatchAnyKeyword(s.PeekAt(0), parse.KeywordCodeStruct):
		s.Move(1)

		s.Skip()
		if !parse.MatchAnySpecial(s.PeekAt(0), '<') {
			return nil, s.WrapError(fmt.Errorf(`invalid struct type field: '<' not found`))
		}
		s.Move(1)

		s.Skip()
		if parse.MatchAnySpecial(s.PeekAt(0), '>') {
			s.Move(1)
			return types.AcceptStruct(s, nil), nil
		}

		var fields []types.StructFieldNode
		for {
			s.Skip()
			field, err := ParseStructField(s.Child())
			if err != nil {
				return nil, s.WrapError(fmt.Errorf(`invalid struct type field: %w`, err))
			}
			s.Move(field.Len())
			fields = append(fields, field)

			s.Skip()
			switch {
			default:
				return nil, s.WrapError(fmt.Errorf(`invalid struct type field: ',' or '>' not found`))
			case parse.MatchAnySpecial(s.PeekAt(0), ','):
				s.Move(1)
			case parse.MatchAnySpecial(s.PeekAt(0), '>'):
				s.Move(1)
				return types.AcceptStruct(s, fields), nil
			}
		}
	}
}

func parseEnclosedTypeSize(s *parse.ParseState) (types.TypeSizeNode, error) {
	s.Skip()
	if !parse.MatchAnySpecial(s.PeekAt(0), '(') {
		return nil, s.WrapError(fmt.Errorf(`invalid type size: '(' not found`))
	}
	s.Move(1)

	s.Skip()
	size, err := ParseTypeSize(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf(`invalid type size: %w`, err))
	}
	s.Move(size.Len())

	s.Skip()
	if !parse.MatchAnySpecial(s.PeekAt(0), ')') {
		return nil, s.WrapError(fmt.Errorf(`invalid type size: ')' not found`))
	}
	s.Move(1)

	return size, nil
}
