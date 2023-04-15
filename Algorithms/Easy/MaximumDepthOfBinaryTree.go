package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MaxIntt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func MaxIntt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
