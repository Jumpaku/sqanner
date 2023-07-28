package parse

import "github.com/Jumpaku/sqanner/tokenize"

type ParseState struct {
	Depth  int
	Cursor int
	Input  []tokenize.Token
}

func (s *ParseState) len() int {
	return len(s.Input[s.Cursor:])
}

func (s *ParseState) peek() tokenize.Token {
	return s.Input[s.Cursor]
}

func (s *ParseState) Up() {
	s.Depth--
}

func (s *ParseState) Down() {
	s.Depth++
}

func (s *ParseState) Next() tokenize.Token {
	t := s.Input[s.Cursor]
	s.Cursor++
	return t
}

func (s *ParseState) SkipSpaces() {
	for s.Expect(func(t tokenize.Token) bool { return t.Kind == tokenize.TokenSpace }) {
		s.Next()
	}
}

func (s *ParseState) Expect(expect func(token tokenize.Token) bool) bool {
	if s.Cursor == s.len() {
		return false
	}
	return expect(s.peek())
}
