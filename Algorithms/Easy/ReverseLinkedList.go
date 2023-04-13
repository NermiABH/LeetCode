package main

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	next := head
	for next != nil {
		curr := newHead
		newHead = next
		next = next.Next
		newHead.Next = curr
	}
	return newHead
}
