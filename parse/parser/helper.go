package paeser

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

func isIdentifier(pattern string) func(t tokenize.Token) bool {
	return func(t tokenize.Token) bool {
		switch t.Kind {
		default:
			return false
		case tokenize.TokenIdentifier:
			return strings.ToLower(pattern) == strings.ToLower(string(t.Content))
		case tokenize.TokenIdentifierQuoted:
			c := t.Content
			return strings.ToLower(pattern) == strings.ToLower(string(c[1:len(c)-1]))
		}
	}
}
