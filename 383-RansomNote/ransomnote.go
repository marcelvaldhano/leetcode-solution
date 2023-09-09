package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {

	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}

	}
	return true

}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote = "aa"
	magazine = "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote = "aa"
	magazine = "aab"
	fmt.Println(canConstruct(ransomNote, magazine))

}
