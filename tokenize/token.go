package tokenize

import "github.com/Jumpaku/go-assert"

// TokenKind represents the type of token identified by the token scanner.
type TokenKind int

const (
	// TokenUnspecified represents an unspecified or unknown token.
	TokenUnspecified TokenKind = iota
	// TokenEOF represents the end of the file (EOF) token.
	TokenEOF
	// TokenSpace represents a space token.
	TokenSpace
	// TokenComment represents a comment token.
	TokenComment
	// TokenIdentifier represents an identifier token.
	TokenIdentifier
	// TokenIdentifierQuoted represents a quoted identifier token.
	TokenIdentifierQuoted
	// TokenLiteralQuoted represents a quoted literal (string) token.
	TokenLiteralQuoted
	// TokenLiteralInteger represents an integer literal token.
	TokenLiteralInteger
	// TokenLiteralFloat represents a floating-point literal token.
	TokenLiteralFloat
	// TokenKeyword represents a keyword token.
	TokenKeyword
	// TokenSpecialChar represents a special character token.
	TokenSpecialChar
)

// The String method for TokenKind returns the name of the token code as a string.
// If an invalid TokenKind is encountered, it will raise an assertion error.
func (k TokenKind) String() string {
	switch k {
	default:
		assert.State(false, `invalid TokenKind: %d`, k)
		return assert.Unexpected1[string](`invalid TokenKind is unexpected`)
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

// Token represents a single token identified by the token scanner.
// The Token struct is used to represent identified tokens during the tokenization process.
// It contains information about the type and location of the token in the source code.
// A valid token has its TokenKind set to a specific type (not TokenUnspecified).
// It contains the following fields:
// - Kind: The TokenKind representing the type of the token.
// - Content: The content or value of the token.
// - Begin: The starting position (index) of the token in the input sequence.
// - End: The ending position (index) of the token in the input sequence.
// - Line: The line number where the token starts in the input source.
// - Column: The column number where the token starts in the input source.
type Token struct {
	Kind    TokenKind
	Content []rune

	Begin  int
	End    int
	Line   int
	Column int
}

// IsValid checks if the token is valid, i.e., its TokenKind is not TokenUnspecified.
// It returns true if the token is valid and false otherwise.
func (t Token) IsValid() bool {
	return t.Kind != TokenUnspecified
}
