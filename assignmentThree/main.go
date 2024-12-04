package main

import (
	"fmt"
	Solutions "github.com/vaidehiadhi/assignmentThree/Solutions"
)

func main() {
	// giving the data
	result := Solutions.Details{}

	fmt.Scanln(&result.Name)
	fmt.Scanln(&result.Age)
	fmt.Scanln(&result.PhoneNumber)
	fmt.Scanln(&result.Address.City)
	fmt.Scanln(&result.Address.State)
	fmt.Scanln(&result.Address.Pincode)

	fmt.Println("Name is", result.Name)
	fmt.Println("Age is", result.Age)
	fmt.Println("Phone number is", result.PhoneNumber)
	fmt.Println("City is", result.Address.City)
	fmt.Println("State is", result.Address.State)
	fmt.Println("Pincode is", result.Address.Pincode)

	var list Solutions.LinkedList

	list.InsertLast(10)
	list.InsertLast(20)
	list.InsertLast(30)
	list.Display()

	list.DeleteLast()
	list.Display()

	list.DeleteLast()
	list.Display()

	list.DeleteLast()
	list.Display()
}
