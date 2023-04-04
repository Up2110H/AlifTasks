package stack

import "errors"

type Stack struct {
	base []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.base) == 0
}

func (s *Stack) Push(val int) {
	s.base = append(s.base, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	last := len(s.base) - 1
	val := s.base[last]

	s.base = s.base[0:last]

	return val, nil
}
