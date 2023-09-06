package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	count := make([]int, 26)

	if len(s) != len(t) {
		return false
	}

	for i, _ := range s {
		count[int(s[i]-'a')] += 1
		count[int(t[i]-'a')] -= 1
	}
	for i, _ := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
