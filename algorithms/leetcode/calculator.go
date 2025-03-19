package main

// Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.

import (
	"strings"
)

func find_end(s string) int {
	counter := 1
	for pos, char := range s {
		switch char {
		case '(':
			counter += 1
		case ')':
			counter -= 1
		}
		if counter == 0 {
			return pos
		}
	}
	return -1
}

func calculate(s string) int {
	res := 0
	sign := 1
	last := 0
	s = strings.Replace(s, " ", "", -1)
	for pos, char := range s {
		switch char {
		case '(':
			{
				end := find_end(s[pos+1:]) + pos + 1
				return res + sign*calculate(s[pos+1:end]) + calculate(s[end+1:])
			}
		case '-':
			{
				sign = -1
				last = 0
			}
		case '+':
			{
				sign = 1
				last = 0
			}
		default:
			{
				x := sign * int(char-'0')
				res += last*9 + x
				last = last*10 + x
			}
		}
	}

	return res
}
