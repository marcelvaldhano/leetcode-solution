package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{}
	}
	start := head
	end := head

	for start.Next != nil && end.Next != nil {
		start = start.Next
		end = end.Next
		if end.Next != nil {
			end = end.Next
		}
	}
	return start
}
