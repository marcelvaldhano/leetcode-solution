package main

import (
	"fmt"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isValid(s string) bool {
	var stack Stack
	p := strings.Split(s, "")

	for i := range p {
		if p[i] == "{" || p[i] == "(" || p[i] == "[" {
			stack.Push(p[i])
			continue
		}
		x, _ := stack.Pop()
		if p[i] == "}" && x != "{" {
			return false
		} else if p[i] == ")" && x != "(" {
			return false
		} else if p[i] == "]" && x != "[" {
			return false
		}

	}
	if stack.IsEmpty() {
		return true
	}
	return false

}

func main() {
	input := "()"
	fmt.Println(isValid(input))
}
