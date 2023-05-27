package main

func inorderTraversalRecursive(root *TreeNode) []int {
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

func inorderTraversalIterative(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack, root = append(stack, root), root.Left
		} else {
			root, stack = stack[len(stack)-1], stack[0:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
