package expr

import (
	"bytes"
	"cloud.google.com/go/civil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jumpaku/sqanner/parse"
	"github.com/Jumpaku/sqanner/parse/node/ddl/expr"
	"github.com/Jumpaku/sqanner/parse/node/types"
	"github.com/Jumpaku/sqanner/tokenize"
	"math/big"
	"strconv"
	"time"
)

func ParseScalar(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	switch {
	default:
		return nil, s.WrapError(fmt.Errorf("fail to parse scalar literals"))
	case parse.MatchAnyKeyword(t, parse.KeywordCodeTrue, parse.KeywordCodeFalse):
		return ParseBool(s.Child())
	case parse.MatchAnyTokenKind(t, tokenize.TokenLiteralInteger):
		return ParseInt64(s.Child())
	case parse.MatchAnyTokenKind(t, tokenize.TokenLiteralFloat):
		return ParseFloat64(s.Child())
	case parse.MatchAnyTokenKind(t, tokenize.TokenLiteralQuoted): // STRING or BYTES
		return parseStringOrBytesLiteral(s.Child())
	case parse.MatchIdentifier(t, true, "DATE"):
		return ParseDate(s.Child())
	case parse.MatchIdentifier(t, true, "JSON"):
		return ParseJSON(s.Child())
	case parse.MatchIdentifier(t, true, "NUMERIC"):
		return ParseNumeric(s.Child())
	case parse.MatchIdentifier(t, true, "TIMESTAMP"):
		return ParseTimestamp(s.Child())
	}
}

func ParseBool(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	switch {
	default:
		return nil, s.WrapError(fmt.Errorf("fail to parse BOOL literal"))
	case parse.MatchAnyKeyword(t, parse.KeywordCodeTrue):
		s.Move(1)
		return expr.AcceptBool(s, "TRUE"), nil
	case parse.MatchAnyKeyword(t, parse.KeywordCodeFalse):
		s.Move(1)
		return expr.AcceptBool(s, "FALSE"), nil
	}
}

func ParseInt64(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenLiteralInteger) {
		return nil, s.WrapError(fmt.Errorf("INT64 literal not found"))
	}

	v := string(t.Content)
	if _, err := strconv.ParseInt(v, 0, 64); err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid INT64 literal: %w", err))
	}

	s.Move(1)
	return expr.AcceptInt64(s, v), nil
}

func ParseFloat64(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenLiteralFloat) {
		return nil, s.WrapError(fmt.Errorf("FLOAT64 literal not found"))
	}

	v := string(t.Content)
	if _, _, err := new(big.Float).Parse(v, 0); err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid FLOAT64 literal: %w", err))
	}

	s.Move(1)
	return expr.AcceptFloat64(s, v), nil
}

func ParseDate(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	if !parse.MatchIdentifier(s.PeekAt(0), true, `DATE`) {
		return nil, s.WrapError(fmt.Errorf("DATE prefix not found"))
	}
	s.Move(1)
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenLiteralQuoted) {
		return nil, s.WrapError(fmt.Errorf("DATE string literal not found"))
	}

	strNode, err := ParseString(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid DATE string literal: %w", err))
	}
	if strNode.Type() != types.TypeCodeString {
		return nil, s.WrapError(fmt.Errorf("DATE string literal not found"))
	}

	v := strNode.UnquotedStringValue()
	if _, err := civil.ParseDate(v); err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid DATE string literal: %w", err))
	}

	s.Move(strNode.Len())

	return expr.AcceptDate(s, v), nil
}

func ParseJSON(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	if !parse.MatchIdentifier(s.PeekAt(0), true, `JSON`) {
		return nil, s.WrapError(fmt.Errorf("JSON prefix not found"))
	}
	s.Move(1)
	s.Skip()

	strNode, err := ParseString(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid JSON string literal: %w", err))
	}
	if strNode.Type() != types.TypeCodeString {
		return nil, s.WrapError(fmt.Errorf("JSON string literal not found"))
	}
	s.Move(strNode.Len())

	v := strNode.UnquotedStringValue()
	var a any
	if err := json.NewDecoder(bytes.NewBufferString(v)).Decode(&a); err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid JSON string literal: %w", err))
	}
	return expr.AcceptJSON(s, v), nil
}

func ParseNumeric(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	if !parse.MatchIdentifier(s.PeekAt(0), true, `NUMERIC`) {
		return nil, s.WrapError(fmt.Errorf("NUMERIC prefix not found"))
	}
	s.Move(1)
	s.Skip()

	strNode, err := ParseString(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid NUMERIC string literal: %w", err))
	}
	if strNode.Type() != types.TypeCodeString {
		return nil, s.WrapError(fmt.Errorf("NUMERIC string literal not found"))
	}
	s.Move(strNode.Len())

	v := strNode.UnquotedStringValue()
	{
		_, e1 := tokenize.ScanNext([]rune(v), tokenize.TokenLiteralFloat)
		_, e2 := tokenize.ScanNext([]rune(v), tokenize.TokenLiteralInteger)
		if e1 != nil && e2 != nil {
			return nil, s.WrapError(fmt.Errorf("invalid NUMERIC string literal: %w", errors.Join(e1, e2)))
		}
	}

	if _, ok := new(big.Rat).SetString(v); !ok {
		return nil, s.WrapError(fmt.Errorf("invalid NUMERIC string literal"))
	}

	return expr.AcceptNumeric(s, v), nil
}

func ParseTimestamp(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	if !parse.MatchIdentifier(s.PeekAt(0), true, `TIMESTAMP`) {
		return nil, s.WrapError(fmt.Errorf("TIMESTAMP prefix not found"))
	}
	s.Move(1)
	s.Skip()

	strNode, err := ParseString(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid TIMESTAMP string literal: %w", err))
	}
	s.Move(strNode.Len())

	v := strNode.UnquotedStringValue()

	if _, err = time.Parse(time.RFC3339, v); err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid TIMESTAMP string literal: %w", err))
	}

	return expr.AcceptTimestamp(s, v), nil
}

func ParseString(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	n, err := parseStringOrBytesLiteral(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid STRING literal: %w", err))
	}
	if n.Type() != types.TypeCodeString {
		return nil, s.WrapError(fmt.Errorf("STRING literal not found"))
	}
	s.Move(n.Len())

	return n, err
}

func ParseBytes(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	n, err := parseStringOrBytesLiteral(s.Child())
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid BYTES literal: %w", err))
	}
	if n.Type() != types.TypeCodeBytes {
		return nil, s.WrapError(fmt.Errorf("BYTES literal not found"))
	}
	s.Move(n.Len())

	return n, err
}

func parseStringOrBytesLiteral(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenLiteralQuoted) {
		return nil, s.WrapError(fmt.Errorf("quoted literal not found"))
	}
	s.Move(1)

	return expr.AcceptStringOrBytes(s, string(t.Content)), nil
}
