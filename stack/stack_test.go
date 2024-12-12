package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		operations []string
		values     []int
		expected   []int
	}{
		{
			operations: []string{"Push", "Push", "Pop", "Pop"},
			values:     []int{1, 2, 0, 0},
			expected:   []int{2, 1},
		},
		{
			operations: []string{"Push", "Pop", "Pop"},
			values:     []int{1, 0, 0},
			expected:   []int{1, 0},
		},
		{
			operations: []string{"Push", "Push", "Pop", "Pop", "Pop"},
			values:     []int{1, 2, 0, 0, 0},
			expected:   []int{2, 1, 0},
		},
	}

	for _, t := range tests {
		stack := &Stack{}
		var result []int
		for i, op := range t.operations {
			if op == "Push" {
				stack.Push(t.values[i])
			} else if op == "Pop" {
				val, _ := stack.Pop()
				result = append(result, val)
			}
		}
	}
}
