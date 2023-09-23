package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, nil, nil)
}

func isValidBST1(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
		return false
	}
	return isValidBST1(root.Left, min, &root.Val) && isValidBST1(root.Right, &root.Val, max)
}
