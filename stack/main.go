package stack

import (
	"fmt"
)

// element struct
type Element struct {
	Value interface{}
}

// Stack struct
type Stack struct {
	Storage []Element
}

// New Stack
func New() *Stack {
	s := Stack{Storage: []Element{}}
	return &s
}

// Push items onto the stack
func (s *Stack) Push(v interface{}) {
	e := Element{Value: v}
	s.Storage = append(s.Storage, e)
}

// Pop items from the stack
func (s *Stack) Pop() interface{} {
	out, storage := s.Storage[len(s.Storage)-1], s.Storage[:len(s.Storage)-1]

	s.Storage = storage
	return out.Value
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.Storage)
}

// Print loops through and prints out the contents of the stack
func (s Stack) Print() {
	for _, e := range s.Storage {
		fmt.Println(e.Value)
	}
}
