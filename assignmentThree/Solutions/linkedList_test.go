package Solutions

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		operations []string
		values     []int
		expected   string
	}{
		{
			operations: []string{"InsertLast", "InsertLast", "Display"},
			values:     []int{1, 2},
			expected:   "1 -> 2 -> nil",
		},
		{
			operations: []string{"InsertLast", "DeleteLast", "Display"},
			values:     []int{1},
			expected:   "nil",
		},
		{
			operations: []string{"InsertLast", "InsertLast", "DeleteLast", "Display"},
			values:     []int{1, 2},
			expected:   "1 -> nil",
		},
	}

	for _, tt := range tests {
		ll := &LinkedList{}
		for i, op := range tt.operations {
			if op == "InsertLast" {
				ll.InsertLast(tt.values[i])
			} else if op == "DeleteLast" {
				ll.DeleteLast()
			}
		}
	}
}
