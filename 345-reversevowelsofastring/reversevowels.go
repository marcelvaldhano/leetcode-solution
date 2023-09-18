package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	split := strings.Split(s, "")
	start := 0
	end := len(split) - 1

	for start < end {
		if !isVowel(split[end]) {
			end--
			continue
		}
		if !isVowel(split[start]) {
			start++
			continue
		}
		split[start], split[end] = split[end], split[start]
		end--
		start++
	}
	return strings.Join(split, "")
}

func isVowel(r string) bool {
	return r == "a" || r == "e" || r == "i" || r == "o" || r == "u" || r == "A" || r == "E" || r == "I" || r == "O" || r == "U"
}
func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
