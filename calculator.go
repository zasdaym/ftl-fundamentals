// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the product of both.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the division of the first number with the second.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide by 0")
	}

	return a / b, nil
}

// Sqrt returns a square root of given number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return -1, fmt.Errorf("invalid number")
	}

	var z float64 = 1
	for i := 1; i < 10; i++ {
		z = (z - (math.Pow(z, 2)-a)/(2*z))
	}

	return z, nil
}
