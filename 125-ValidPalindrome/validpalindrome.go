package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern := "[^a-zA-Z0-9]+"

	re := regexp.MustCompile(pattern).ReplaceAllString(s, "")

	split := strings.Split(re, " ")
	s = strings.Join(split, "")
	s = strings.ToLower(s)

	start := 0
	lengthString := len(s) - 1

	for start < lengthString {
		if s[start] != s[lengthString] {
			return false
		}
		start++
		lengthString--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s = "race a car"
	fmt.Println(isPalindrome(s))
	s = " "
	fmt.Println(isPalindrome(s))
}
