package main

import (
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	val, ok := stack.Pop()
	if ok {
		fmt.Printf("Popped value: %d\n", val)
	} else {
		fmt.Println("Stack is empty")
	}
}
