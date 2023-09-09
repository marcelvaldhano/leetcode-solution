package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	res := strings.Trim(s, " ")
	split := strings.Split(res, " ")
	return len(split[len(split)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))

}
