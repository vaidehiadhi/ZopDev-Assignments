package main

import (
	"fmt"
	"math"
)

// creating interface
type Shape interface {
	Area() float64 //methods
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {

	var s Shape
	s = Circle{5}
	fmt.Printf("area of cirle is :%v\n", s.Area())
	s = Rectangle{7, 4}
	fmt.Printf("area of rectangle :%v\n", s.Area())

}
