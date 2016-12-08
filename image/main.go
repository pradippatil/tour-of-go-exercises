//Solution to excercise: https://tour.golang.org/methods/25
package main

import "golang.org/x/tour/pic"
import "image/color"
import "image"

type Image struct {
	Height, Width int
}

// ColorModel returns the Image's color model.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel

}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Height, img.Width)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 155, 155}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)

	// put resulting base64 string in this converter to see image
	// http://codebeautify.org/base64-to-image-converter
}
