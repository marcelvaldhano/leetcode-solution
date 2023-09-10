package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(split) != len(pattern) {
		return false
	}
	dictPattern, dictS := map[string]string{}, map[string]string{}
	for i, v := range pattern {
		value, value2 := dictPattern[string(v)], dictS[split[i]]
		if value == "" && value2 == "" {
			dictPattern[string(v)], dictS[split[i]] = split[i], string(s)
			continue
		}
		if value != split[i] && string(v) != value2 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat fish"))
}
