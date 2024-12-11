package main

import (
	"fmt"
)

// struct
type Calculator struct {
	result int
}

// method to add
func (c *Calculator) Add(a int, b int) int {
	return a + b
}

// method to sub
func (c *Calculator) Sub(a int, b int) int {
	return a - b
}

// methods to mul
func (c *Calculator) Mul(a int, b int) int {
	return a * b
}

// method to div
func (c *Calculator) Div(a int, b int) int {
	return a / b
}

func main() {

	//initalizing struct
	cal := Calculator{0}

	//calling methods
	fmt.Printf("the addition is: %v\n", cal.Add(5, 4))
	fmt.Printf("the subtraction is: %v\n", cal.Sub(7, 2))
	fmt.Printf("the multiplication is: %v\n", cal.Mul(5, 10))
	fmt.Printf("the Division is: %v\n", cal.Div(60, 10))

}
