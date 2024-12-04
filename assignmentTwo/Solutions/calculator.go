package Solutions

import "fmt"

func Calulator() {
	fmt.Println("enter the first number")
	fmt.Println("enter the second number")
	fmt.Println("enter the operator")
	var val1, val2 float64
	var operator string
	fmt.Scanln(&val1)
	fmt.Scanln(&val2)
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f + %f = %f\n", val1, val2, val1+val2)
	case "-":
		fmt.Printf("%f - %f = %f\n", val1, val2, val1-val2)
	case "*":
		fmt.Printf("%f * %f = %f\n", val1, val2, val1*val2)
	case "/":
		fmt.Printf("%f / %f = %f\n", val1, val2, val1/val2)
	default:
		fmt.Println("invalid choice")
	}
}
