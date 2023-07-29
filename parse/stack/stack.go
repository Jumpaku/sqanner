package stack

type Stack[T any] struct {
	data []T
	len  int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() T {
	return s.data[s.len-1]
}

func (s *Stack[T]) Push(e T) *Stack[T] {
	if len(s.data) == s.len {
		s.data = append(s.data, e)
	} else {
		s.data[s.len] = e
	}
	s.len++

	return s
}

func (s *Stack[T]) Pop() T {
	s.len--
	return s.data[s.len]
}

func (s *Stack[T]) Len() int {
	return s.len
}
