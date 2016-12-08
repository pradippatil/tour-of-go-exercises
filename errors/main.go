// Soultion to excercise https://tour.golang.org/methods/20
package main

import (
	"fmt"
	"math"
	"os"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function calculates square root of a given positive number using Netwon's approximation method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, d := 1.0, 1.0
	for math.Abs(d) > 1e-15 {
		zPrevious := z
		z = z - ((z*z - x) / (2 * z))
		d = zPrevious - z
		fmt.Printf("Current delta: %v, z: %v, zPrevious: %v\n", d, z, zPrevious)
	}
	return z, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}

func main() {
	value, err := Sqrt(-4)
	if err != nil {
		checkError(err)
	}
	fmt.Printf("Using Newton's method: %v\n", value)
	fmt.Printf("Using builtin function: %v\n", math.Sqrt(4))

	// TODO
	// Try this formula as well
	// For number N, and initial approximation as 1
	// newApproximation =  (oldApproximation + N/oldApproximation)/2
	// http://www.dave-reed.com/book/Chapter8/Newton.html
	// More details : https://math.mit.edu/~stevenj/18.335/newton-sqrt.pdf
}
