package main

func postorderTraversalRecursive(root *TreeNode) []int {
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

func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, result := make([]*TreeNode, 0), make([]int, 0)
	stack = append(stack, root)
	for {
		if root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
				root = root.Left
			} else {
				root = root.Right
			}
		} else {
			root, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		Loop:
			result = append(result, root.Val)
			if len(stack) == 0 {
				return result
			}
			peek := stack[len(stack)-1]
			if peek.Left == root || peek.Right == root {
				stack = stack[0 : len(stack)-1]
				root = peek
				goto Loop
			}
			root = peek
		}
	}
}
