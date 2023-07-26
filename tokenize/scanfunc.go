package tokenize

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
	"unicode"
)

// Spaces scans the input sequence represented by the ScanState 's' to find the number of consecutive space runes at the current Cursor position.
// It returns the count of runes in the scanned space token and the corresponding TokenCode.
// If no spaces are found at the current Cursor position, the function returns 0 for the count and TokenUnspecified for the TokenCode.
// If an error occurs during processing, it will be returned as the third value, which will be nil in this implementation.
func Spaces(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 || !unicode.IsSpace(s.PeekAt(0)) {
		return 0, TokenUnspecified, nil
	}

	n := s.CountWhile(0, unicode.IsSpace)

	return n, TokenSpace, nil
}

// Comment scans the input sequence represented by the ScanState 's' to identify and handle comments.
// It returns the count of runes in the scanned comment token and the corresponding TokenCode.
// If no comments are found at the current Cursor position, the function returns 0 for the count and TokenUnspecified for the TokenCode.
// If the comment starts with '#' and extends to the end of the line, the function returns the count of runes up to the newline character.
// If the comment starts with '//' or '--' and extends to the end of the line, the function returns the count of runes up to the newline character.
// If the comment starts with '/*' and ends with '*/', the function returns the count of runes up to the closing '*/' sequence.
// If the comment is not properly terminated with '*/', the function returns an error with a message indicating an incomplete comment.
func Comment(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 {
		return 0, TokenUnspecified, nil
	}
	if s.PeekAt(0) == '#' {
		foundAt, found := s.FindFirst(1, 1, func(r []rune) bool { return string(r) == "\n" })
		if found {
			return foundAt + 1, TokenComment, nil
		}
		return s.Len(), TokenComment, nil
	}

	if s.Len() < 2 {
		return 0, TokenUnspecified, nil
	}
	switch string(s.PeekSlice(0, 2)) {
	default:
		return 0, TokenUnspecified, nil
	case `//`, `--`:
		foundAt, found := s.FindFirst(2, 1, func(r []rune) bool { return string(r) == "\n" })
		if found {
			return foundAt + 1, TokenComment, nil
		}
		return s.Len(), TokenComment, nil
	case `/*`:
		foundAt, found := s.FindFirst(2, 2, func(r []rune) bool { return string(r) == `*/` })
		if !found {
			return 0, TokenUnspecified, fmt.Errorf(`comment incompleted: "*/" is expected but not found`)
		}
		return foundAt + 2, TokenComment, nil
	}
}

// expectEscapeSequence returns:
//   - (0, nil) if s.PeekAt(cur) starts with not backslash
//   - (<runes of escape sequence>, nil) if valid escape sequence is scanned
//   - (0, err) if invalid escape sequence is detected
//
// Escape sequence: https://cloud.google.com/spanner/docs/reference/standard-sql/lexical#escape_sequences
func expectEscapeSequence(cur int, s *ScanState) (int, error) {
	errInvalidEscape := func(es string) error { return fmt.Errorf(`invalid excape sequence: %q`, es) }
	if cur < s.Len() && string(s.PeekAt(cur)) != `\` {
		return 0, nil
	}

	if cur+2 > s.Len() {
		return 0, errInvalidEscape(string(s.PeekSlice(cur, s.Len())))
	}

	switch s.PeekAt(cur + 1) {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '?', '"', '\'', '`':
		size := 2

		return size, nil
	case '0', '1', '2', '3', '4', '5', '6', '7':
		size := 4
		if cur+size > s.Len() {
			return 0, errInvalidEscape(string(s.PeekSlice(cur, s.Len())))
		}
		if v := s.PeekSlice(cur, cur+size); slices.ContainsFunc(v[1:], func(r rune) bool { return !isOctalDigit(r) }) {
			return 0, errInvalidEscape(string(v))
		}

		return size, nil
	case 'x', 'X':
		size := 4
		if cur+size > s.Len() {
			return 0, errInvalidEscape(string(s.PeekSlice(cur, s.Len())))
		}
		if v := s.PeekSlice(cur, cur+size); slices.ContainsFunc(v[2:], func(r rune) bool { return !isHexDigit(r) }) {
			return 0, errInvalidEscape(string(v))
		}

		return size, nil
	case 'u':
		size := 6
		if cur+size > s.Len() {
			return 0, errInvalidEscape(string(s.PeekSlice(cur, s.Len())))
		}
		if v := s.PeekSlice(cur, cur+size); slices.ContainsFunc(v[2:], func(r rune) bool { return !isHexDigit(r) }) {
			return 0, errInvalidEscape(string(v))
		}

		return size, nil
	case 'U':
		size := 10
		if cur+size > s.Len() {
			return 0, errInvalidEscape(string(s.PeekSlice(cur, s.Len())))
		}
		if v := s.PeekSlice(cur, cur+size); slices.ContainsFunc(v[2:], func(r rune) bool { return !isHexDigit(r) }) {
			return 0, errInvalidEscape(string(v))
		}

		return size, nil
	default:
		return 0, errInvalidEscape(string(s.PeekAt(cur)))
	}
}
func expectQuotedToken(quote []rune, prefix []rune, escape bool, s *ScanState) (int, error) {
	cur := len(prefix)
	for {
		foundAt, found := s.FindFirst(cur, len(quote), func(r []rune) bool {
			return string(r) == string(quote) || (escape && r[0] == '\\')
		})
		if !found {
			return 0, fmt.Errorf("incomplete quoted token: closing quote %q is expected but not found", quote)
		}
		switch s.PeekAt(foundAt) {
		case '\\': // Escape sequence
			size, err := expectEscapeSequence(foundAt, s)
			if err != nil {
				return 0, fmt.Errorf(`quoted token including invalid excape sequence: %w`, err)
			}

			cur += size
		default: // Close quotation
			cur = foundAt
			return cur + len(quote), nil
		}
	}
}

// IdentifierQuoted scans the input sequence represented by the ScanState 's' to identify and handle quoted identifiers enclosed within back quotes (`).
// It returns the count of runes in the scanned quoted identifier token and the corresponding TokenCode.
// If no quoted identifier is found at the current Cursor position, the function returns 0 for the count and TokenUnspecified for the TokenCode.
// If the quoted identifier is empty (two consecutive backticks), the function returns an error indicating an empty quoted identifier.
// If the quoted identifier is not properly enclosed within backticks, the function returns an error with a message indicating an invalid quoted identifier.
func IdentifierQuoted(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 || s.PeekAt(0) != '`' {
		return 0, TokenUnspecified, nil
	}

	if s.Len() >= 2 && s.PeekAt(1) == '`' {
		return 0, TokenUnspecified, fmt.Errorf("quorted identifier cannot be empty")
	}

	size, err := expectQuotedToken([]rune("`"), []rune("`"), true, s)
	if err != nil {
		return 0, TokenUnspecified, fmt.Errorf(`invalid quorted identifier: %w`, err)
	}
	return size, TokenIdentifierQuoted, nil
}

// LiteralQuoted scans the input sequence represented by the ScanState 's' to identify and handle quoted literals (strings or bytes) with optional prefixes.
// It returns the count of runes in the scanned quoted literal, the corresponding TokenCode, and an error if any occurs during processing.
// If no quoted literal is found at the current Cursor position, the function returns 0 for the count, TokenUnspecified for the TokenCode, and nil for the error.
func LiteralQuoted(s *ScanState) (int, TokenCode, error) {
	errInvalidQuotedLiteral := func(err error) error { return fmt.Errorf(`invalid quorted literal: %w`, err) }

	getPrefix := func() []rune {
		for _, br := range []string{``, `b`, `r`, `br`, `rb`} {
			for _, quote := range []string{`"""`, `'''`, `"`, `'` /* """ and ''' must precede " and ' */} {
				brQuotes := []rune(br + quote)
				n := len(brQuotes)
				if s.Len() < n {
					continue
				}

				if prefix := s.PeekSlice(0, n); strings.ToLower(string(prefix)) == string(brQuotes) {
					return prefix
				}
			}
		}
		return nil
	}

	prefix := getPrefix()
	escape := !(slices.Contains(prefix, 'r') || slices.Contains(prefix, 'R'))
	switch strings.ToLower(string(prefix)) {
	default:
		return 0, TokenUnspecified, nil
	case `"`, `b"`, `r"`, `rb"`, `br"`, `'`, `b'`, `r'`, `rb'`, `br'`:
		size, err := expectQuotedToken(prefix[len(prefix)-1:], prefix, escape, s)
		if err != nil {
			return 0, TokenUnspecified, errInvalidQuotedLiteral(err)
		}
		return size, TokenLiteralQuoted, nil
	case `"""`, `b"""`, `r"""`, `rb"""`, `br"""`, `'''`, `b'''`, `r'''`, `rb'''`, `br'''`:
		size, err := expectQuotedToken(prefix[len(prefix)-3:], prefix, escape, s)
		if err != nil {
			return 0, TokenUnspecified, errInvalidQuotedLiteral(err)
		}
		return size, TokenLiteralQuoted, nil
	}
}

// IdentifierOrKeyword scans the input sequence represented by the ScanState 's' to identify and handle identifiers or keywords.
// It returns the count of runes in the scanned identifier or keyword, the corresponding TokenCode, and an error if any occurs during processing.
// If no identifier or keyword is found at the current Cursor position, the function returns 0 for the count, TokenUnspecified for the TokenCode, and nil for the error.
// If the scanned token is a keyword, the function returns the count of runes in the scanned keyword and TokenCode TokenKeyword.
// If the scanned token is an identifier, the function returns the count of runes in the scanned identifier and TokenCode TokenIdentifier.
func IdentifierOrKeyword(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 || !isLetter(s.PeekAt(0)) {
		return 0, TokenUnspecified, nil
	}

	n := s.CountWhile(1, func(r rune) bool { return isLetter(r) || isDecimalDigit(r) })
	size := n + 1
	if isKeyword(s.PeekSlice(0, size)) {
		return size, TokenKeyword, nil
	}
	return size, TokenIdentifier, nil
}

// NumberOrDot scans the input sequence represented by the ScanState 's' to identify and handle numbers or the dot (.) operator.
// It returns the count of runes in the scanned number or dot operator, the corresponding TokenCode, and an error if any occurs during processing.
// If no number or dot operator is found at the current Cursor position, the function returns 0 for the count, TokenUnspecified for the TokenCode, and nil for the error.
// The function recognizes hexadecimal integers (starting with "0x"), decimals (with or without a decimal point), and floating-point numbers (with or without an exponent using 'e' or 'E').
// If the scanned token is the dot (.) operator, the function returns 1 for the count and TokenCode TokenSpecialChar.
// If the scanned token is an integer (either decimal or hexadecimal), the function returns the count of runes in the scanned integer and TokenCode TokenLiteralInteger.
// If the scanned token is a floating-point number, the function returns the count of runes in the scanned number and TokenCode TokenLiteralFloat.
func NumberOrDot(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 || !(isDecimalDigit(s.PeekAt(0)) || s.PeekAt(0) == '.') {
		return 0, TokenUnspecified, nil
	}

	// hexadecimal
	if s.Len() >= 2 && strings.ToLower(string(s.PeekSlice(0, 2))) == `0x` {
		nDigits := s.CountWhile(2, isHexDigit)
		if nDigits == 0 {
			return 0, TokenUnspecified, fmt.Errorf(`incomplete hex integer: hex digits not found after %s`, string(s.PeekSlice(0, 2)))
		}

		return nDigits + 2, TokenLiteralInteger, nil
	}

	// dot operator
	if s.PeekAt(0) == '.' && s.CountWhile(1, isDecimalDigit) == 0 {
		return 1, TokenSpecialChar, nil
	}

	cur := 0
	if isDecimalDigit(s.PeekAt(cur)) {
		cur += s.CountWhile(cur, isDecimalDigit)
		// decimal
		if cur == s.Len() || !(s.PeekAt(cur) == '.' || s.PeekAt(cur) == 'e' || s.PeekAt(cur) == 'E') {
			return cur, TokenLiteralInteger, nil
		}
	}

	// float
	if s.PeekAt(cur) == '.' {
		cur++
		cur += s.CountWhile(cur, isDecimalDigit)
	}

	if cur == s.Len() || !(s.PeekAt(cur) == 'e' || s.PeekAt(cur) == 'E') {
		return cur, TokenLiteralFloat, nil
	}

	// exponential
	cur++
	if cur == s.Len() {
		return 0, TokenUnspecified, fmt.Errorf(`incomplete float: [+-] or digits not found after %s`, string(s.PeekSlice(0, cur)))
	}

	if cur < s.Len() && (s.PeekAt(cur) == '-' || s.PeekAt(cur) == '+') {
		cur++
	}

	nExp := s.CountWhile(cur, isDecimalDigit)
	if nExp == 0 {
		return 0, TokenUnspecified, fmt.Errorf(`incomplete float: digits not found after %s`, string(s.PeekSlice(0, cur)))
	}

	return cur + nExp, TokenLiteralFloat, nil
}

// SpecialChar scans the input sequence represented by the ScanState 's' to identify and handle special characters.
// It returns the count of runes in the scanned special character, the corresponding TokenCode, and an error if any occurs during processing.
// If no special character is found at the current Cursor position, the function returns 0 for the count, TokenUnspecified for the TokenCode, and nil for the error.
// If the scanned token is a dot (.) character followed by a decimal digit, the function returns 0 for the count, TokenUnspecified for the TokenCode, and nil.
// For all other cases, where the current rune represents a standalone special character, the function returns 1 for the count and TokenCode TokenSpecialChar.
func SpecialChar(s *ScanState) (int, TokenCode, error) {
	if s.Len() == 0 || !isSpecialChar(s.PeekAt(0)) {
		return 0, TokenUnspecified, nil
	}

	if s.Len() >= 2 && s.PeekAt(0) == '.' && isDecimalDigit(s.PeekAt(1)) {
		return 0, TokenUnspecified, nil
	}

	return 1, TokenSpecialChar, nil
}
