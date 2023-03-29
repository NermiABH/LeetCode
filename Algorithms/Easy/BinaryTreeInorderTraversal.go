package main

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ans)
	*ans = append(*ans, root.Val)
	helper(root.Right, ans)
}
