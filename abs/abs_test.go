package abs

import (
	"testing"
)

// declaring a Test function
func TestAbs(t *testing.T) {

	//defining a struct
	test := []struct {
		desp      string
		input     int
		expOutput int
	}{ //initializing the struct with test cases
		{"positive number", 1, 1},
		{"negative number", -1, 1},
	}

	//looping over the struct
	for _, v := range test {
		//calling the function
		res := Abs(v.input)
		//checking
		if res != v.expOutput {
			t.Errorf("Test failed: %s. Input: %d, Expected: %d, Got: %d",
				v.desp, v.input, v.expOutput, res)

		}

	}
}
