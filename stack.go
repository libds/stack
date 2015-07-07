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

// Push down an item to the stack
func (s *Stack) Push(v string) {
	s.items = append([]string{v}, s.items...)
}

// Return and remove the first item in the list
func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", StackEmpty
	}
	v := s.items[0]
	s.items = append(s.items[:0], s.items[1:]...)
	return v, nil
}

// Returns true if stack is empty
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

// Returns the amount of items currently in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Get the underlying array of items
func (s *Stack) Arr() []string {
	return s.items
}
