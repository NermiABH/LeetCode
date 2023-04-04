package main

func postorderTraversal(root *TreeNode) []int {
	var slice []int
	var helper func(root *TreeNode, slice *[]int)
	helper = func(root *TreeNode, slice *[]int) {
		if root == nil {
			return
		}
		helper(root.Left, slice)
		helper(root.Right, slice)
		*slice = append(*slice, root.Val)
	}
	helper(root, &slice)
	return slice
}
