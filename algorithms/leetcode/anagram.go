package main

import "reflect"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise

func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, c := range s {
		sMap[c] += 1
	}

	for _, c := range t {
		tMap[c] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}
