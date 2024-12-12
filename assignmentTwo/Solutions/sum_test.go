package Solutions

import (
	"testing"
)

func TestSum(t *testing.T) {
	test := []struct {
		desp   string
		input  int
		output int
	}{
		{"correct sum", 2, 3},
		//{"incorrect sum", 2, 2},
	}

	for _, v := range test {
		res := Sum(v.input)
		if res != v.output {
			t.Errorf("invalid: %v, input: %v, output: %v", v.desp, v.input, v.output)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	{
		test := []struct {
			desp   string
			input  int
			output int
		}{
			{"correct sum", 2, 3},
			//{"incorrect sum", 2, 2},
		}
		for range b.N {
			for _, v := range test {
				Sum(v.input)
			}

		}
	}
}
