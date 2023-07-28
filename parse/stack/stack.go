package stack

type Stack[T any] struct {
	data []T
	top  int
}

func (s *Stack[T]) Top() T {
	return s.data[s.top]
}

func (s *Stack[T]) Push(e T) *Stack[T] {
	if len(s.data) == s.top {
		s.data = append(s.data, e)
	} else {
		s.data[s.top] = e
	}
	s.top++

	return s
}

func (s *Stack[T]) Pop() T {
	s.top--
	return s.data[s.top+1]
}
