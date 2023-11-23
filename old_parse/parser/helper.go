package parser

import (
	"github.com/Jumpaku/sqanner/old_parse/node"
	"github.com/Jumpaku/sqanner/tokenize"
	"golang.org/x/exp/slices"
	"strings"
)

func IsAnyKind(t tokenize.Token, k ...tokenize.TokenKind) bool {
	return slices.Contains(k, t.Kind)
}
func IsAnySpecial(t tokenize.Token, r ...rune) bool {
	return t.Kind == tokenize.TokenSpecialChar && slices.Contains(r, t.Content[0])
}

func IsKeyword(t tokenize.Token, code node.KeywordCode) bool {
	return t.Kind == tokenize.TokenKeyword && node.KeywordCodeOf(string(t.Content)) == code
}

func IsIdentifier(t tokenize.Token, ignoreCase bool, pattern ...string) bool {
	if t.Kind != tokenize.TokenIdentifier && t.Kind != tokenize.TokenIdentifierQuoted {
		return false
	}

	if ignoreCase {
		q := strings.ToLower(string(t.Content))
		return slices.ContainsFunc(pattern, func(p string) bool { return q == strings.ToLower(p) })
	}

	return slices.Contains(pattern, string(t.Content))
}
func isIdentifier(ignoreCase bool, pattern ...string) func(t tokenize.Token) bool {
	return func(t tokenize.Token) bool {
		if t.Kind != tokenize.TokenIdentifier && t.Kind != tokenize.TokenIdentifierQuoted {
			return false
		}

		if ignoreCase {
			q := strings.ToLower(string(t.Content))
			return slices.ContainsFunc(pattern, func(p string) bool { return q == strings.ToLower(p) })
		}

		return slices.Contains(pattern, string(t.Content))
	}
}
