package main

import (
	"fmt"
	"math"
)

// Newton's Method
func Sqrt(x float64) float64 {
	var z = 1.
	for i := 0; i < 10; i++ {
		// fmt.Println(z)
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func Sqrt(x float64) float64 {
	var z = z - (z*z-x)/(2*z)
	for old := x; new-z > .00000001 {
		// fmt.Println(z)
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	num := 401.
	fmt.Println("Finding the square root of", num)
	fmt.Println("Newton =", Sqrt(num))
	fmt.Println("actual =", math.Sqrt(num))
}
