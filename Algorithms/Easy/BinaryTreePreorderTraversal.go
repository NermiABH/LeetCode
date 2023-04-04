package main

func preorderTraversal(root *TreeNode) []int {
	var slice []int
	var helper func(root *TreeNode, slice *[]int)
	helper = func(root *TreeNode, slice *[]int) {
		if root == nil {
			return
		}
		*slice = append(*slice, root.Val)
		helper(root.Left, slice)
		helper(root.Right, slice)
	}
	helper(root, &slice)
	return slice
}
