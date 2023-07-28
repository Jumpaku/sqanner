package tokenize

import "github.com/Jumpaku/go-assert"

// TokenCode represents the type of token identified by the token scanner.
// The following constants define different token codes that the scanner can return:
// - TokenUnspecified: Represents an unspecified or unknown token.
// - TokenEOF: Represents the end of the file (EOF) token.
// - TokenSpace: Represents a space token.
// - TokenComment: Represents a comment token.
// - TokenIdentifier: Represents an identifier token.
// - TokenIdentifierQuoted: Represents a quoted identifier token.
// - TokenLiteralQuoted: Represents a quoted literal (string) token.
// - TokenLiteralInteger: Represents an integer literal token.
// - TokenLiteralFloat: Represents a floating-point literal token.
// - TokenKeyword: Represents a keyword token.
// - TokenSpecialChar: Represents a special character token.
type TokenCode int

const (
	// TokenUnspecified represents an unspecified or unknown token.
	TokenUnspecified TokenCode = iota
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

// The String method for TokenCode returns the name of the token code as a string.
// If an invalid TokenCode is encountered, it will raise an assertion error.
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

// Token represents a single token identified by the token scanner.
// The Token struct is used to represent identified tokens during the tokenization process.
// It contains information about the type and location of the token in the source code.
// A valid token has its TokenCode set to a specific type (not TokenUnspecified).
// It contains the following fields:
// - Code: The TokenCode representing the type of the token.
// - Content: The content or value of the token.
// - Begin: The starting position (index) of the token in the input sequence.
// - End: The ending position (index) of the token in the input sequence.
// - Line: The line number where the token starts in the input source.
// - Column: The column number where the token starts in the input source.
type Token struct {
	Code    TokenCode
	Content []rune

	Begin  int
	End    int
	Line   int
	Column int
}

// IsValid checks if the token is valid, i.e., its TokenCode is not TokenUnspecified.
// It returns true if the token is valid and false otherwise.
func (t Token) IsValid() bool {
	return t.Code != TokenUnspecified
}
