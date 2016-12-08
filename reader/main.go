//Solution to excercise: https://tour.golang.org/methods/22
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// populate given byte array with 'A'
func (m MyReader) Read(p []byte) (int, error) {
	p[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
