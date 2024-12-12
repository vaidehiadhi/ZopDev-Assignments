package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		operation string
		a, b      int
		expected  int
	}{
		{
			operation: "Add",
			a:         5,
			b:         4,
			expected:  9,
		},
		{
			operation: "Sub",
			a:         7,
			b:         2,
			expected:  5,
		},
		{
			operation: "Mul",
			a:         5,
			b:         10,
			expected:  50,
		},
		{
			operation: "Div",
			a:         60,
			b:         10,
			expected:  6,
		},
	}

	cal := &Calculator{0}

	for _, v := range tests {
		var result int
		switch v.operation {
		case "Add":
			result = cal.Add(v.a, v.b)
		case "Sub":
			result = cal.Sub(v.a, v.b)
		case "Mul":
			result = cal.Mul(v.a, v.b)
		case "Div":
			result = cal.Div(v.a, v.b)
		}

		if result != v.expected {
			t.Errorf("For operation %s with inputs %d and %d, expected %d, got %d", v.operation, v.a, v.b, v.expected, result)
		}

	}
}
