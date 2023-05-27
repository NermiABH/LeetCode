package main

func preorderTraversalRecursive(root *TreeNode) []int {
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

func preorderTraversalIterative(root *TreeNode) []int {
	stack, result := make([]*TreeNode, 0), make([]int, 0)
	for {
		if root != nil {
			result = append(result, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		} else {
			if len(stack) == 0 {
				return result
			}
			root, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		}
	}
}
