package Solutions

import "fmt"

func Sum(number int) int {
	var sum int
	for i := 0; i <= number; i++ {
		sum += i
		fmt.Printf("%d + %d = %d\n", number, i, sum)
	}
	return sum
}
