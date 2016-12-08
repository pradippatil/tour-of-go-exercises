// Solution to excercise https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

// Pic returns a dx X dy matrix with random uint8 integers which are used to render base64 string
// Can use any base64 to img converter to check resulting image
// e.g. this online converter http://codebeautify.org/base64-to-image-converter#
func Pic(dx, dy int) [][]uint8 {

	matrix := make([][]uint8, dx)
	for i := range matrix {
		matrix[i] = make([]uint8, dy)
		for j := range matrix[i] {
			matrix[i][j] = uint8(i + j)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
