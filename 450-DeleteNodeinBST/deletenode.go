package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		temp := findMin(root.Right)
		root.Val = temp.Val
		root.Right = deleteNode(root.Right, temp.Val)
	}
	return root
}
func findMin(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return findMin(node.Left)
}
