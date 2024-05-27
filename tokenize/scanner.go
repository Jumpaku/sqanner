package tokenize

import (
	"fmt"
)

// ScanState represents a state for scanning and processing a sequence of runes.
// It contains the Input slice, which holds the runes to be scanned, and the Cursor
// indicating the current position in the Input slice.
type ScanState struct {
	Input  []rune
	Cursor int
}

// Len returns the remaining number of runes in the Input slice from the current Cursor position.
func (s ScanState) Len() int {
	return len(s.Input) - s.Cursor
}

// PeekAt returns the rune at a given relative position 'offset' from the current Cursor position.
// 'offset' must be in [0, s.Len()).
func (s ScanState) PeekAt(offset int) rune {
	return s.Input[s.Cursor+offset]
}

// PeekSlice returns a slice of runes starting from 'begin' to 'endExclusive' positions from the current Cursor position.
// 'begin' must be in [0, s.Len()).
// 'endExclusive' must be in [0, s.Len()].
func (s ScanState) PeekSlice(begin int, endExclusive int) []rune {
	return s.Input[s.Cursor+begin : s.Cursor+endExclusive]
}

// CountWhile counts the number of runes that satisfy the given function 'satisfy', starting from the 'begin' position from the current Cursor position.
// The counting stops as soon as a rune that does not satisfy the condition is encountered.
// 'begin' must be in [0, s.Len()).
func (s ScanState) CountWhile(begin int, satisfy func(rune) bool) int {
	count := 0
	for n := begin; n < s.Len(); n++ {
		if satisfy(s.Input[s.Cursor+n]) {
			count++
		} else {
			break
		}
	}
	return count
}

// FindFirst searches for the first occurrence of a pattern with the given 'patternSize' using the 'pattern' function, starting from the 'begin' position from the current Cursor position.
// It returns the index of the first occurrence and a boolean indicating if the pattern was found.
// If the pattern is not found, it returns the index 's.Len()' and 'false'.
// 'begin' must be in [0, s.Len()).
func (s ScanState) FindFirst(begin int, patternSize int, pattern func([]rune) bool) (int, bool) {
	for n := begin; n < s.Len()-patternSize+1; n++ {
		if pattern(s.Input[s.Cursor+n : s.Cursor+n+patternSize]) {
			return n, true
		}
	}

	return s.Len(), false
}

// TokenScanner provides a tokenizer for processing a sequence of runes and identifying different types of tokens.
type TokenScanner struct {
	ScanState
	lines   int
	columns int
}

func (s *TokenScanner) Init(input []rune) {
	s.ScanState = ScanState{Input: input}
}

// ScanNext scans the next token in the input sequence and returns the Token and an error if any occurs during processing.
// If the end of the input sequence is reached, the method returns a special Token with TokenKind TokenEOF to indicate the end of the file.
func (s *TokenScanner) ScanNext() (Token, error) {
	if s.Len() == 0 {
		return s.accept(0, TokenEOF), nil
	}

	type scanFunc func(s *ScanState) (int, TokenKind, error)
	var scanners = []scanFunc{
		/* scanFunc in front have higher priority than those behind. */
		Spaces,
		Comment,
		IdentifierQuoted,
		LiteralQuoted,
		IdentifierOrKeyword,
		NumberOrDot,
		SpecialChar,
	}

	for _, scanner := range scanners {
		size, kind, err := scanner(&s.ScanState)
		if err != nil {
			return Token{}, s.wrapErr(err)
		}
		if kind != TokenUnspecified {
			return s.accept(size, kind), nil
		}
	}

	return Token{}, s.wrapErr(fmt.Errorf(`invalid character sequense`))
}

func (s *TokenScanner) accept(n int, kind TokenKind) Token {
	out := s.Input[s.Cursor : s.Cursor+n]
	token := Token{
		Kind:    kind,
		Content: out,
		Begin:   s.Cursor,
		End:     s.Cursor + n,
		Line:    s.lines,
		Column:  s.columns,
	}
	for _, o := range out {
		if o == '\n' {
			s.lines++
			s.columns = 0
		}
		s.columns++
		s.Cursor++
	}
	return token
}

func (s *TokenScanner) wrapErr(err error) error {
	sizeAfter := 20
	if sizeAfter > s.Len() {
		sizeAfter = s.Len()
	}
	sizeBefore := 5
	if sizeBefore > s.Cursor {
		sizeBefore = s.Cursor
	}
	input := string(s.Input[s.Cursor-sizeBefore : s.Cursor+sizeAfter])
	return fmt.Errorf(`fail to scan token at line %d column %d near ...%q...: %w`, s.lines, s.columns, input, err)
}

// Tokenize returns a slice of Token representing the identified tokens in the input sequence.
// If an error occurs during tokenization, the function returns an error with a message indicating the failure to tokenize.
func Tokenize(input []rune) ([]Token, error) {
	scanner := &TokenScanner{}
	scanner.Init(input)
	var tokens []Token
	for {
		token, err := scanner.ScanNext()
		if err != nil {
			return nil, fmt.Errorf(`fail to tokenize: %w`, err)
		}
		tokens = append(tokens, token)
		if token.Kind == TokenEOF {
			return tokens, nil
		}
	}
}

func ScanNext(input []rune, k TokenKind) (token []rune, err error) {
	s := &ScanState{Input: input}
	switch k {
	default:
		return nil, fmt.Errorf("invalid token kind specified: %v", k)
	case TokenEOF:
		if s.Len() != 0 {
			return nil, fmt.Errorf("%v not found", k)
		}
		return []rune{}, nil
	case TokenSpace:
		n, _, err := Spaces(s)
		if err != nil {
			return nil, err
		}
		return s.Input[:n], nil
	case TokenComment:
		n, _, err := Comment(s)
		if err != nil {
			return nil, err
		}
		return s.Input[:n], nil
	case TokenIdentifier, TokenKeyword:
		n, kind, err := IdentifierOrKeyword(s)
		if err != nil {
			return nil, err
		}
		if k != kind {
			return nil, fmt.Errorf("%v not found", k)
		}
		return s.Input[:n], nil
	case TokenIdentifierQuoted:
		n, _, err := IdentifierQuoted(s)
		if err != nil {
			return nil, err
		}
		return s.Input[:n], nil
	case TokenLiteralQuoted:
		n, _, err := LiteralQuoted(s)
		if err != nil {
			return nil, err
		}
		return s.Input[:n], nil
	case TokenLiteralInteger, TokenLiteralFloat:
		n, kind, err := NumberOrDot(s)
		if err != nil {
			return nil, err
		}
		if k != kind {
			return nil, fmt.Errorf("%v not found", k)
		}
		return s.Input[:n], nil
	case TokenSpecialChar:
		n, kind, err := SpecialChar(s)
		if err != nil {
			return nil, err
		}
		if k != kind {
			return nil, fmt.Errorf("%v not found", k)
		}
		return s.Input[:n], nil
	}
}
