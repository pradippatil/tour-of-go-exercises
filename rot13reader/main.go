// Solution to excercise:https://tour.golang.org/methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt *rot13Reader) Read(p []byte) (int, error) {
	n, err := rt.r.Read(p) //read using io.Reader

	// rotate each alphabet byte by 13 positions
	// https://en.wikipedia.org/wiki/ROT13
	for i := range p {
		switch {
		case (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm'):
			p[i] += 13
		case (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z'):
			p[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
