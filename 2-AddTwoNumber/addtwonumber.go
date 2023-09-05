package main

import "fmt"

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		carry = temp.Val / 10
		temp.Val = temp.Val % 10
		if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
		}
		if carry > 0 {
			temp.Next = &ListNode{Val: carry}
		}
		temp = temp.Next
	}
	return result
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	printLinkedList(addTwoNumbers(&l1, &l2))

}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}
