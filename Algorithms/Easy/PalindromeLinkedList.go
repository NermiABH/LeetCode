package main

func isPalindromeListNode(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	slow = reverseList(slow)
	for slow.Next != nil {
		if head.Val != slow.Val {
			return false
		}
		head, slow = head.Next, slow.Next
	}
	return true
}

//func reverseList(head *ListNode) *ListNode {
//	var newHead *ListNode
//	next := head
//	for next != nil {
//		curr := newHead
//		newHead = next
//		next = next.Next
//		newHead.Next = curr
//	}
//	return newHead
//}
