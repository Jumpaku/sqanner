package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
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

func (s *ParseState) Child() *ParseState {
	return &ParseState{
		input:  s.input,
		begin:  s.cursor,
		cursor: s.cursor,
	}
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

func (s *ParseState) SkipSpacesAndComments() {
	offset := 0
	for offset < s.Len() && IsAnyKind(s.PeekAt(offset), tokenize.TokenComment, tokenize.TokenSpace) {
		offset++
	}
	s.Move(offset)
}

func Error(s *ParseState, err error) error {
	t := s.PeekAt(0)

	begin := s.cursor
	if begin > 0 {
		begin--
	}

	end := s.cursor + 1
	if s.Len() > 0 {
		end++
	}

	var contents []string
	for _, token := range s.input[begin:end] {
		contents = append(contents, string(token.Content))
	}

	return fmt.Errorf(`fail to parse during processing tokens near ...%v...: line=%d, column=%d: %w`, contents, t.Line, t.Column, err)
}

func Accept[T node.Node](s *ParseState, nodeFunc node.NewNodeFunc[T]) T {
	return nodeFunc(s.begin, s.cursor)
}
