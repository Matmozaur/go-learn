// What is the most common word in Rumi's poem?
package main

import (
	"fmt"
	"sort"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
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
	// fmt.Println("a"[0])

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

	/*
	   Print character and frequency in percent of Rumi's poem
	   Print in frequency descending order
	   Assume ASCII text, ignore white space
	*/

	charFrequency := make(map[rune]int, 27)
	sum := 0
	// fmt.Println(string(rune(97)))
	for _, c := range poem {
		if (c <= 122) && (c >= 97) {
			charFrequency[c] += 1
			sum += 1
		}
	}

	fmt.Println("Char frequency:")
	for c, f := range charFrequency {
		fmt.Println(fmt.Sprintf(" %s: %d", string(c), f))
	}

	// sorting
	keys := make([]rune, 0, len(charFrequency))
	for key := range charFrequency {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return charFrequency[keys[i]] < charFrequency[keys[j]]
	})

	charFrequencySorted := orderedmap.New[string, float32]()
	for _, c := range keys {
		charFrequencySorted.Set(string(c), float32(charFrequency[c])/float32(sum))
	}

	fmt.Println("Char frequency sorted:")
	for pair := charFrequencySorted.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%s: %f\n", pair.Key, pair.Value)
	}
}
