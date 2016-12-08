// Soultion to excercise https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

// Sqrt function calculates square root of a given positive number using Netwon's approximation method
func Sqrt(x float64) float64 {
	z, d := 1.0, 1.0
	for math.Abs(d) > 1e-15 {
		zPrevious := z
		z = z - ((z*z - x) / (2 * z))
		d = zPrevious - z
		fmt.Printf("Current delta: %v, z: %v, zPrevious: %v\n", d, z, zPrevious)
	}
	return z
}

func main() {
	fmt.Printf("Using Newton's method: %v\n", Sqrt(4))
	fmt.Printf("Using builtin function: %v\n", math.Sqrt(4))

	// TODO
	// Try this formula as well
	// For number N, and initial approximation as 1
	// newApproximation =  (oldApproximation + N/oldApproximation)/2
	// http://www.dave-reed.com/book/Chapter8/Newton.html
	// More details : https://math.mit.edu/~stevenj/18.335/newton-sqrt.pdf
}
