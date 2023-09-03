package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := -121
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	delimiter := ""
	str := strconv.Itoa(x)
	split := strings.Split(str, delimiter)
	i := 0
	j := len(split) - 1
	for i < j {
		if split[i] != split[j] {
			return false
		}
		i++
		j--
	}
	return true
}
