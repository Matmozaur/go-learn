// What is the most common word in Rumi's poem?
package main

import (
	"fmt"
	"strings"
)

var poem = `
those who do not feel this love
pulling them like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let them sleep
`

func main() {
	frequency := make(map[string]int)
	fmt.Println(frequency)
	for _, word := range strings.Fields(poem) {
		frequency[word] += 1
	}
	fmt.Println(frequency)

	// maxW, maxC := "", 0
	// for w, c := range frequency {
	// 	if c > maxC {
	// 		maxW, maxC = w, c
	// 	}
	// }

	// fmt.Println(maxW, maxC)
	maxWs, maxC := make([]string, 0), 0
	for w, c := range frequency {
		if c > maxC {
			maxWs, maxC = []string{w}, c
		} else if c == maxC {
			maxWs = append(maxWs, w)
		}
	}
	fmt.Println(maxWs, maxC)
}
