package Solutions

import (
	"testing"
)

func TestGreet(t *testing.T) {
	test := []struct {
		desp    string
		valid   string
		invalid string
	}{
		{"this is a valid string", "vaidehi", "vaidehi"},
		{"this is not a valid string", "hello", "vaidehi"},
	}

	for _, v := range test {
		res := Greet()
		if res != v.valid {
			t.Errorf("Test failed: %s, valid: %s, invalid: %s", v.desp, v.valid, v.invalid)
		}
	}
}
