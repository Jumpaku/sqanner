package parse

import (
	"fmt"
	"github.com/Jumpaku/sqanner/tokenize"
)

type ParseState struct {
	input  []tokenize.Token
	begin  int
	cursor int
}

func NewParseState(input []tokenize.Token) *ParseState {
	return &ParseState{input: input}
}

func (s *ParseState) clone() *ParseState {
	return &ParseState{
		input:  s.input,
		begin:  s.begin,
		cursor: s.cursor,
	}
}
func (s *ParseState) Child() *ParseState {
	child := s.clone()
	child.begin = s.cursor
	return child
}

func (s *ParseState) PeekAt(offset int) tokenize.Token {
	return s.input[s.cursor+offset]
}

func (s *ParseState) Len() int {
	return len(s.input) - s.cursor
}

func (s *ParseState) Move(offset int) {
	s.cursor += offset
}

func (s *ParseState) Skip() {
	offset := 0
	for offset < s.Len() && MatchAnyTokenKind(s.PeekAt(offset), tokenize.TokenComment, tokenize.TokenSpace) {
		offset++
	}
	s.Move(offset)
}
func (s *ParseState) Begin() int {
	return s.begin
}
func (s *ParseState) End() int {
	return s.cursor
}

func (s *ParseState) WrapError(err error) error {
	t := s.PeekAt(0)

	begin := s.cursor - 1
	if begin < 0 {
		begin = 0
	}

	end := s.cursor + 2
	if end > len(s.input) {
		end = len(s.input)
	}

	var contents []string
	for _, token := range s.input[begin:end] {
		contents = append(contents, string(token.Content))
	}

	return fmt.Errorf(`fail to parse during processing tokens near ...%v...: line=%d, column=%d: %w`, contents, t.Line, t.Column, err)
}
