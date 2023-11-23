package parse

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"slices"
	"strings"
)

func MatchAnyTokenKind(t tokenize.Token, k ...tokenize.TokenKind) bool {
	return slices.Contains(k, t.Kind)
}

func MatchIdentifier(t tokenize.Token, ignoreCase bool, pattern ...string) bool {
	if t.Kind != tokenize.TokenIdentifier && t.Kind != tokenize.TokenIdentifierQuoted {
		return false
	}

	if ignoreCase {
		q := strings.ToLower(string(t.Content))
		return slices.ContainsFunc(pattern, func(p string) bool { return q == strings.ToLower(p) })
	}

	return slices.Contains(pattern, string(t.Content))
}

func MatchAnySpecial(t tokenize.Token, r ...rune) bool {
	return t.Kind == tokenize.TokenSpecialChar && slices.Contains(r, t.Content[0])
}

func MatchAnyKeyword(t tokenize.Token, codes ...KeywordCode) bool {
	k, err := OfKeyword(string(t.Content))
	if err != nil {
		return false
	}
	return slices.ContainsFunc(codes, func(code KeywordCode) bool { return k == code })
}
