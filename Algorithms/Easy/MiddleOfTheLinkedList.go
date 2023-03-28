package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// middleNode2 second solution of this problem
func middleNode2(head *ListNode) *ListNode {
	length := getLength(head)
	ans := head
	for i := 0; i < length/2; i++ {
		ans = ans.Next
	}
	return ans
}
func getLength(head *ListNode) int {
	next, i := head.Next, 1
	for next != nil {
		next = next.Next
		i++
	}
	return i
}
