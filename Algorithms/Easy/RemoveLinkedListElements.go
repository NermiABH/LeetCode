package main

func removeElements(head *ListNode, val int) *ListNode {
	current := &ListNode{Next: head}
	head = current
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head.Next
}
