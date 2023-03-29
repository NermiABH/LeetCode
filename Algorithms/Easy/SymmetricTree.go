package main

func isSymmetric(root *TreeNode) bool {
	var helper func(left *TreeNode, right *TreeNode) bool
	helper = func(left *TreeNode, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		} else if left.Val != right.Val {
			return false
		}
		return helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}
	return helper(root.Left, root.Right)
}
