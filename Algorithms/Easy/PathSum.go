package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	} else if targetSum = targetSum - root.Val; (targetSum == 0 && root.Left == nil && root.Right == nil) ||
		hasPathSum(root.Left, targetSum) ||
		hasPathSum(root.Right, targetSum) {
		return true
	}
	return false
}
