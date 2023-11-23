package stack

type Stack[T any] struct {
	head T
	tail *Stack[T]
	len  int
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s Stack[T]) Peek() T {
	return s.head
}

func (s Stack[T]) Push(e T) Stack[T] {
	return Stack[T]{head: e, tail: &s, len: s.len + 1}
}

func (s Stack[T]) Pop() Stack[T] {
	return *s.tail
}

func (s Stack[T]) Len() int {
	return s.len
}
