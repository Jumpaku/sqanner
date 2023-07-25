package tokenize

import (
	"fmt"
)

type ScanState struct {
	Input  []rune
	Cursor int
}

func (s ScanState) Len() int {
	return len(s.Input) - s.Cursor
}
func (s ScanState) PeekAt(n int) rune {
	return s.Input[s.Cursor+n]
}
func (s ScanState) PeekSlice(begin int, endExclusive int) []rune {
	return s.Input[s.Cursor+begin : s.Cursor+endExclusive]
}
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
func (s ScanState) FindFirst(begin int, patternSize int, pattern func([]rune) bool) (int, bool) {
	for n := begin; n < s.Len()-patternSize+1; n++ {
		if pattern(s.Input[s.Cursor+n : s.Cursor+n+patternSize]) {
			return n, true
		}
	}

	return s.Len(), false
}

type TokenScanner struct {
	ScanState
	lines   int
	columns int
}

func NewTokenScanner(input []rune) *TokenScanner {
	return &TokenScanner{ScanState: ScanState{Input: input}}
}

func (s *TokenScanner) ScanNext() (Token, error) {
	if s.Len() == 0 {
		return s.accept(0, TokenEOF), nil
	}

	type scanFunc func(s *TokenScanner) (int, TokenCode, error)
	var scanners = []scanFunc{
		Spaces,
		Comment,
		IdentifierQuoted,
		LiteralQuoted,
	}
	for _, scanner := range scanners {
		n, code, err := scanner(s)
		if err != nil {
			return Token{}, s.wrapErr(err)
		}
		if code != TokenUnspecified {
			return s.accept(n, code), nil
		}
	}
	return Token{}, s.wrapErr(fmt.Errorf(`invalid character sequense`))
	/*
		switch {
		case isDigit(s.PeekAt(0)):
			n := s.CountWhile(0, unicode.IsDigit)
			if s.PeekAt(n) == '.' {
			}
			m := s.countIf(func(r rune) bool { return strings.ContainsRune(digitChars, r) })
			if s.PeekAt(n) == '.' {
			}
		case s.PeekAt(0) == '.':
			if strings.ContainsRune(digitChars, s.PeekAt(1)) {
			}
			return s.accept(0, TokenUnspecified), nil
		case strings.ContainsRune(specialChars, s.PeekAt(0)):
			return s.accept(0, TokenUnspecified), nil
		case strings.ContainsRune(letterChars, s.PeekAt(0)):
			n := s.countIf(func(r rune) bool { return strings.ContainsRune(letterChars+digitChars, r) })
			lowerSlice := strings.ToUpper(string(s.PeekSlice(0, n)))
			if keywords[lowerSlice] {
				return s.accept(0, TokenUnspecified), nil
			}
			return s.accept(0, TokenUnspecified), nil
		}
		return s.accept(0, TokenUnspecified), nil
	*/
}

func (s *TokenScanner) accept(n int, code TokenCode) Token {
	out := s.Input[s.Cursor : s.Cursor+n]
	token := Token{
		Code:    code,
		Content: string(out),
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
	return fmt.Errorf(`fail to scan token at line %d column %d near ...%s...: %w`, s.lines, s.columns, input, err)
}
