//Solution to excercise : https://tour.golang.org/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b, count := 0, 1, -1
	return func() int {
		if count == -1 || count == 0 {
			count++
			return count
		}
		sum := a + b
		a = b
		b = sum
		return sum
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
