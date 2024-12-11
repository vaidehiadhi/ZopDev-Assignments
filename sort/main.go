package main

import (
	"fmt"
)

func sort(s []int, compare func(a int, b int) bool) {
	for i := 0; i < len(s); i++ {
		minInd := i

		for j := i + 1; j < len(s); j++ {
			if compare(s[j], s[minInd]) {
				minInd = j
			}
		}
		s[i], s[minInd] = s[minInd], s[i]

	}
}

func compare(a int, b int) bool {
	if a < b {
		return true
	}
	return false
}

func main() {
	s := []int{6, 2, 1, 9, 0}
	sort(s, compare)
	fmt.Printf("The sorted slice: %v", s)
}
