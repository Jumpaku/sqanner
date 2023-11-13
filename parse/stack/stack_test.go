package stack_test

import (
	"github.com/Jumpaku/sqanner/parse/stack"
	"testing"
)

func check(t *testing.T, got stack.Stack[int], wantLen int, wantPeek int) {
	t.Helper()

	if got.Len() != wantLen {
		t.Errorf(`got stack have size %d but want len is %d`, got.Len(), wantLen)
	}
	if got.Len() == 0 {
		return
	}
	if got.Peek() != wantPeek {
		t.Errorf(`got stack have peek %d but want peek is %d`, got.Peek(), wantPeek)
	}
}

func TestStack(t *testing.T) {
	s := stack.New[int]()
	check(t, s, 0, 0)
	{
		s = s.Push(1)
		check(t, s, 1, 1)
	}
	{
		s = s.Push(2)
		check(t, s, 2, 2)
	}
	{
		s = s.Push(3)
		check(t, s, 3, 3)
	}
	{
		r := s.Peek()
		s = s.Pop()
		check(t, s, 2, 2)
		if r != 3 {
			t.Errorf(`%d is popped but want %d`, r, 3)
		}
	}
	{
		r := s.Peek()
		s = s.Pop()
		check(t, s, 1, 1)
		if r != 2 {
			t.Errorf(`%d is popped but want %d`, r, 2)
		}
	}
	{
		r := s.Peek()
		s = s.Pop()
		check(t, s, 0, 0)
		if r != 1 {
			t.Errorf(`%d is popped but want %d`, r, 1)
		}
	}
	{
		s = s.Push(1)
		check(t, s, 1, 1)
	}
	{
		s = s.Push(2)
		check(t, s, 2, 2)
	}
	{
		s = s.Push(3)
		check(t, s, 3, 3)
	}
}
