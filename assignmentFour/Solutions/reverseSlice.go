package Solutions

import "fmt"

func Reverse() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice is", s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println("The reversed Slice is", s)

}
