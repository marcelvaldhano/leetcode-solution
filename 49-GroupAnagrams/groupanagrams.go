package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	dict := [][]string{
		{strs[0]},
	}
	for i := 1; i < len(strs); i++ {
		for j, value := range dict {
			if isAnagram(string(strs[i]), value[0]) {
				dict[j] = append(dict[j], strs[i])
				break
			}
			if j == len(dict)-1 {
				dict = append(dict, []string{strs[i]})
				break
			}
		}
	}
	return dict
}

func isAnagram(s string, t string) bool {
	count := make([]int, 26)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		count[int(s[i]-'a')] += 1
		count[int(t[i]-'a')] -= 1
	}
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
