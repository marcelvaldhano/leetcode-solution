package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			ok := isPalindrome(s[i : j+1])
			if ok && len(s[i:j+1]) > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))

}
