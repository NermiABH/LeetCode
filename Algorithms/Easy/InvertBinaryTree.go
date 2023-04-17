package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Left), invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
