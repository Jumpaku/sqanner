package parser

import (
	"fmt"
	"github.com/Jumpaku/sqanner/parse/node"
	"github.com/Jumpaku/sqanner/parse/stack"
	"github.com/Jumpaku/sqanner/tokenize"
)

type ParseState struct {
	Cursor int
	begins *stack.Stack[int]
	input  []tokenize.Token
}

func NewParseState(input []tokenize.Token) *ParseState {
	return &ParseState{
		input:  input,
		begins: &stack.Stack[int]{},
	}
}

func Init(s *ParseState) {
	s.begins.Push(s.Cursor)
}

func Accept[T node.Node](s *ParseState, newNode func(int, []tokenize.Token) T) (T, error) {
	head := s.begins.Pop()
	return newNode(head, s.input[head:s.Cursor]), nil
}

func Error[T node.Node](s *ParseState, err error) (T, error) {
	t := s.peek()

	begin := s.Cursor
	if begin > 0 {
		begin--
	}

	end := s.Cursor + 1
	if end < s.len() {
		end++
	}

	var contents []string
	for _, token := range s.input[begin:end] {
		contents = append(contents, string(token.Content))
	}

	var n T
	return n, fmt.Errorf(`fail to parse during processing tokens near ...%v...: line=%d, column=%d: %w`, contents, t.Line, t.Column, err)
}

func (s *ParseState) ExpectNext(expect func(token tokenize.Token) bool) bool {
	if s.Cursor == s.len() {
		return false
	}
	return expect(s.peek())
}

func (s *ParseState) Next() tokenize.Token {
	s.Cursor++
	return s.input[s.Cursor-1]
}

func (s *ParseState) SkipSpacesAndComments() {
	for s.ExpectNext(IsAnyKind(tokenize.TokenSpace, tokenize.TokenComment)) {
		s.Next()
	}
}

func (s *ParseState) len() int {
	return len(s.input[s.Cursor:])
}

func (s *ParseState) peek() tokenize.Token {
	return s.input[s.Cursor]
}
