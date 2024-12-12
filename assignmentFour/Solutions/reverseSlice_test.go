package Solutions

import (
	"testing"
)

func TestReverse(t *testing.T) {

	test := []struct {
		desp   string
		input  []int
		output []int
	}{
		{"the reversed slice", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"the reversed slice", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, v := range test {
		res := Reverse()
		if len(res) != len(v.output) {
			t.Errorf("slice is reversed: %v, input: %v, output: %v", v.desp, v.input, v.output)
			continue
		}

		for i := range res {
			if res[i] != v.output[i] {
				t.Errorf("%s: Mismatch at index %d. Input: %v, Expected: %v, Got: %v",
					v.desp, i, v.input, v.output, res)
				break
			}
		}
	}
}
