package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dict[s[j:i]] && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
