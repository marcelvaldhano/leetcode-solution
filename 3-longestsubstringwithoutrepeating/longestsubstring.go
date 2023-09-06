package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	split := strings.Split(s, "")
	if len(split) == 0 {
		return 0
	}
	if len(split) == 1 {
		return 1
	}
	maks := 0
	result := split[0]
	for i := 1; i < len(split); i++ {
		index, contain := checkContain(result, split[i])
		if contain {
			if maks < len(result) {
				maks = len(result)
			}
			result = result[index+1:]
		}
		result += split[i]
		if i == len(split)-1 && maks < len(result) {
			maks = len(result)
		}
	}
	return maks
}

func checkContain(result string, s string) (int, bool) {
	var index int
	var contain bool
	split := strings.Split(result, "")
	for i, _ := range split {
		if split[i] == s {
			contain = true
			index = i
			break
		}
	}
	return index, contain
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
