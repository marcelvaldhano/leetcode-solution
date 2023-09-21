package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{3, nil, nil}}))
}
