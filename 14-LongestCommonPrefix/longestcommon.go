package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for index, prefix := range strs[0] {
		for _, v := range strs[1:] {
			if index == len(v) || byte(prefix) != v[index] {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
