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
	"regexp"
	"strconv"
	"strings"
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
		return expr.AcceptBool(s, true), nil
	case parse.MatchAnyKeyword(t, parse.KeywordCodeFalse):
		s.Move(1)
		return expr.AcceptBool(s, false), nil
	}
}

func ParseInt64(s *parse.ParseState) (expr.ScalarNode, error) {
	s.Skip()

	t := s.PeekAt(0)
	if !parse.MatchAnyTokenKind(t, tokenize.TokenLiteralInteger) {
		return nil, s.WrapError(fmt.Errorf("INT64 literal not found"))
	}

	v, err := strconv.ParseInt(string(t.Content), 0, 64)
	if err != nil {
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

	f, _, err := new(big.Float).Parse(string(t.Content), 0)
	if err != nil {
		return nil, s.WrapError(fmt.Errorf("invalid FLOAT64 literal: %w", err))
	}

	v, _ := f.Float64()
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

	v, err := civil.ParseDate(strNode.StringValue())
	if err != nil {
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

	dec := json.NewDecoder(bytes.NewBufferString(strNode.StringValue()))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
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

	str := strNode.StringValue()
	{
		_, e1 := tokenize.ScanNext([]rune(str), tokenize.TokenLiteralFloat)
		_, e2 := tokenize.ScanNext([]rune(str), tokenize.TokenLiteralInteger)
		if e1 != nil && e2 != nil {
			return nil, s.WrapError(fmt.Errorf("invalid NUMERIC string literal: %w", errors.Join(e1, e2)))
		}
	}

	v, ok := new(big.Rat).SetString(str)
	if !ok {
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

	v, err := time.Parse(time.RFC3339, strNode.StringValue())
	if err != nil {
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

	var (
		quote   string
		content = string(t.Content)
	)

	prefix := regexp.MustCompile(`^[rbRB]?[rbRB]?`).FindString(content)
	isBytes := strings.ContainsAny(prefix, `bB`)
	isRaw := strings.ContainsAny(prefix, `rR`)
	content = strings.TrimPrefix(content, prefix)

	switch {
	default:
		return nil, s.WrapError(fmt.Errorf("invalid quoted literal"))
	case regexp.MustCompile(`""".*"""$`).MatchString(content):
		quote = `"""`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`'''.*'''$`).MatchString(content):
		quote = `'''`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`".*"$`).MatchString(content):
		quote = `"`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	case regexp.MustCompile(`'.*'$`).MatchString(content):
		quote = `'`
		content = strings.TrimPrefix(strings.TrimSuffix(content, quote), quote)
	}

	if !isRaw {
		unescape := map[string]string{
			`\a`:  "\a",
			`\b`:  "\b",
			`\f`:  "\f",
			`\n`:  "\n",
			`\r`:  "\r",
			`\t`:  "\t",
			`\v`:  "\v",
			`\\`:  "\\",
			`\?`:  "?",
			`\"`:  "\"",
			`\'`:  "'",
			"\\`": "`",
		}
		for from, to := range unescape {
			content = strings.ReplaceAll(content, from, to)
		}

		content = regexp.MustCompile(`\\[0-7]{3}`).ReplaceAllStringFunc(content, unescapeOctal)
		content = regexp.MustCompile(`\\u[0-7a-fA-F]{4}`).ReplaceAllStringFunc(content, unescapeUnicode4)
		content = regexp.MustCompile(`\\U[0-7a-fA-F]{8}`).ReplaceAllStringFunc(content, unescapeUnicode8)
	}

	if isBytes {
		return expr.AcceptBytes(s, []byte(content)), nil
	} else {
		return expr.AcceptString(s, content), nil
	}
}

func unescapeOctal(escaped string) string {
	runes := []rune(escaped)
	o := 0
	for i := 0; i < 3; i++ {
		o *= 8
		o += int(runes[1+i] - '0')
	}
	return string(rune('\000' + o))
}

func unescapeUnicode4(escaped string) string {
	runes := []rune(escaped)
	u := 0
	for i := 0; i < 4; i++ {
		u *= 16
		r := runes[2+i]
		switch {
		case '0' <= r && r <= '9':
			u += int(r - '0')
		case 'a' <= r && r <= 'f':
			u += int(r - 'a')
		case 'A' <= r && r <= 'F':
			u += int(r - 'A')
		}
	}
	return string(rune('\u0000' + u))
}

func unescapeUnicode8(escaped string) string {
	runes := []rune(escaped)
	u := 0
	for i := 0; i < 8; i++ {
		u *= 16
		r := runes[2+i]
		switch {
		case '0' <= r && r <= '9':
			u += int(r - '0')
		case 'a' <= r && r <= 'f':
			u += int(r - 'a')
		case 'A' <= r && r <= 'F':
			u += int(r - 'A')
		}
	}
	return string(rune('\U00000000' + u))
}
