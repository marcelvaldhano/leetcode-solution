package main

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	counter := getLinkedListLength(head) - 1
	result := 0
	for head != nil {
		result += int(math.Pow(2, float64(counter))) * head.Val
		counter--
		head = head.Next
	}
	return result
}

func getLinkedListLength(head *ListNode) int {
	length := 0
	current := head

	for current != nil {
		length++
		current = current.Next
	}

	return length
}

func main() {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	fmt.Println(getDecimalValue(&l1))
}
