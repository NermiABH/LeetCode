func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	recursion(root1, &list1)
	recursion(root2, &list2)
	return reflect.DeepEqual(list1, list2)
}

func recursion(node *TreeNode, leafList *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leafList = append(*leafList, node.Val)
		return
	}

	recursion(node.Left, leafList)
	recursion(node.Right, leafList)
}