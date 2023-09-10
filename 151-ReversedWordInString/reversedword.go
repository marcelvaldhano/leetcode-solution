package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	split := strings.Split(s, " ")

	i, j := 0, len(split)-1

	for i < j {
		split[i], split[j] = split[j], split[i]
		i, j = i+1, j-1
	}
	return strings.Join(split, " ")

}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
