// Solution to excericise : https://tour.golang.org/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

//WordCount returns count of occurance for each word in given string s
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int)
	for i := range words {
		wc[words[i]]++
	}
	return wc

}

func main() {
	wc.Test(WordCount)
}
