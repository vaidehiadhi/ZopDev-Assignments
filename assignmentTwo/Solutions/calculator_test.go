package Solutions

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	test := []struct {
		a      int
		b      int
		op     string
		output int
	}{
		{1, 2, "+", 3},
		{2, 1, "-", 1},
		{2, 3, "*", 6},
		{6, 3, "/", 2},
	}

	for _, v := range test {
		res := Calulator(v.a, v.b, v.op)
		if res != v.output {
			t.Errorf("val 1: %v, val 2: %v, operator: %v, output: %v", v.a, v.b, v.op, v.output)
		}
	}

}

func BenchmarkRandInt(b *testing.B) {

	test := []struct {
		a      int
		b      int
		op     string
		output int
	}{
		{1, 2, "+", 3},
		{2, 1, "-", 1},
		{2, 3, "*", 6},
		{6, 3, "/", 2},
	}
	for range b.N {
		for _, v1 := range test {
			Calulator(v1.a, v1.b, v1.op)
		}
	}
}
