package parse

import (
	"github.com/Jumpaku/sqanner/tokenize"
	"slices"
)

func IsAnyKind(t tokenize.Token, k ...tokenize.TokenKind) bool {
	return slices.Contains(k, t.Kind)
}
