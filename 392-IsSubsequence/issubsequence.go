package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	index := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[index] {
			index++
		}
		if index == len(s) {
			return true
		}
	}

	if index != len(s) {
		return false
	}
	return true

}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s, t = "axc", "ahbgdc"
	fmt.Println(isSubsequence(s, t))

}
