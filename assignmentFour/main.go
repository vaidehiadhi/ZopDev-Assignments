package main

import (
	"fmt"
	"github.com/vaidehiadhi/assignmentFour/Solutions"
)

func main() {
	//ReverseSlice
	Solutions.Reverse()

	//returnMap
	word := "hello"
	result := Solutions.CountCharacters(word)

	fmt.Println("Character counts:", result)

	//sumValue
	m := map[string][]int{
		"a": {1, 2, 3},
		"b": {4, 5},
		"c": {6},
	}

	result = Solutions.SumValuesByKey(m)
	fmt.Println("Summed values by key:", result)

	//SliceToMap
	slice := []string{"apple", "banana", "cherry"}

	resultMap := Solutions.SliceToMap(slice)

	fmt.Println(resultMap)

	//sets
	mySet := Solutions.NewSet()

	// Adding elements
	mySet.Add("Apple")
	mySet.Add("Banana")
	mySet.Add("Cherry")
	mySet.Add("Apple")
	fmt.Println(mySet)

	//Deleting elements
	//mySet.Delete("Apple")

}
