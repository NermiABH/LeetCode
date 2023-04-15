package main

func isBalanced(root *TreeNode) bool {
	return treeDepth(root) != -1
}
func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeDepth(root.Left)
	if left == -1 {
		return -1
	}
	right := treeDepth(root.Right)
	if right == -1 || Abs(left-right) > 1 {
		return -1
	}
	return MaxInt(left, right) + 1
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
