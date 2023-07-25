package tokenize

import "github.com/Jumpaku/go-assert"

type TokenCode int

const (
	TokenUnspecified TokenCode = iota
	TokenEOF
	TokenSpace
	TokenComment
	TokenIdentifier
	TokenIdentifierQuoted
	TokenLiteralQuoted
	TokenLiteralInteger
	TokenLiteralFloat
	TokenKeyword
	TokenSpecialChar
)

func (c TokenCode) String() string {
	switch c {
	default:
		assert.State(false, `invalid TokenCode: %d`, c)
		return assert.Unexpected1[string](`invalid TokenCode is unexpected`)
	case TokenUnspecified:
		return "TokenUnspecified"
	case TokenEOF:
		return "TokenEOF"
	case TokenSpace:
		return "TokenSpace"
	case TokenComment:
		return "TokenComment"
	case TokenIdentifier:
		return "TokenIdentifier"
	case TokenIdentifierQuoted:
		return "TokenIdentifierQuoted"
	case TokenLiteralQuoted:
		return "TokenLiteralQuoted"
	case TokenLiteralInteger:
		return "TokenLiteralInteger"
	case TokenLiteralFloat:
		return "TokenLiteralFloat"
	case TokenKeyword:
		return "TokenKeyword"
	case TokenSpecialChar:
		return "TokenSpecialChar"
	}
}

type Token struct {
	Code    TokenCode
	Content string

	Begin  int
	End    int
	Line   int
	Column int
}

func (t Token) IsValid() bool {
	return t.Code != TokenUnspecified
}
