package main

import (
	"fmt"
)

type Temperature struct {
	result int
}

func (c *Temperature) CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func (c *Temperature) FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	cal := &Temperature{}

	celsius := 25.0
	fahrenheit := cal.CelsiusToFahrenheit(celsius)
	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", celsius, fahrenheit)

	fahrenheit = 77.0
	celsius = cal.FahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", fahrenheit, celsius)
}
