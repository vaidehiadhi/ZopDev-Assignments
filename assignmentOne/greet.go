package assignmentOne

import "fmt"

func Greet() string {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	return "Hello " + name + "!"
}
