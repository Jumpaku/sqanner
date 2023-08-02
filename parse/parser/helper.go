package parser

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"golang.org/x/exp/slices"
	"strings"
)

func IsAnyKind(k ...tokenize.TokenKind) func(t tokenize.Token) bool {
	return func(t tokenize.Token) bool {
		return slices.Contains(k, t.Kind)
	}
}
func isSpecial(r rune) func(t tokenize.Token) bool {
	return func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenSpecialChar && r == t.Content[0]
	}
}

func isKeyword(keyword string) func(t tokenize.Token) bool {
	return func(t tokenize.Token) bool {
		return t.Kind == tokenize.TokenKeyword && strings.ToLower(keyword) == strings.ToLower(string(t.Content))
	}
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
