package main

import (
	"testing"
)

func TestTemperatureConversion(t *testing.T) {
	tests := []struct {
		operation string
		input     float64
		expected  float64
	}{
		{
			operation: "CelsiusToFahrenheit",
			input:     0,
			expected:  32.0,
		},
		{
			operation: "CelsiusToFahrenheit",
			input:     100,
			expected:  212.0,
		},
		{
			operation: "FahrenheitToCelsius",
			input:     32,
			expected:  0.0,
		},
		{
			operation: "FahrenheitToCelsius",
			input:     212,
			expected:  100.0,
		},
	}

	cal := &Temperature{}

	for _, tt := range tests {
		var result float64
		switch tt.operation {
		case "CelsiusToFahrenheit":
			result = cal.CelsiusToFahrenheit(tt.input)
		case "FahrenheitToCelsius":
			result = cal.FahrenheitToCelsius(tt.input)
		}

		if result != tt.expected {
			t.Errorf("For operation %s with input %.2f, expected %.2f, got %.2f", tt.operation, tt.input, tt.expected, result)
		}
	}
}
