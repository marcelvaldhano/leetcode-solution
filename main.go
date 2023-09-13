package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	// Create a map to store the words in wordDict for faster lookups.
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Create a slice to keep track of valid substrings.
	valid := make([]bool, len(s)+1)
	valid[len(s)] = true // An empty string is always valid.

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s); j > i; j-- {
			if valid[j] && wordSet[s[i:j]] {
				valid[i] = true
				break
			}
		}
	}

	return valid[0]
}

// catsandog

func main() {
	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
