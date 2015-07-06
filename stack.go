package stack

import (
	"errors"
)

var StackEmpty = errors.New("Stack is empty")

type Stack struct {
	items []string
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v string) {
	s.items = append([]string{v}, s.items...)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", StackEmpty
	}
	v := s.items[0]
	s.items = append(s.items[:0], s.items[1:]...)
	return v, nil
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Arr() []string {
	return s.items
}
