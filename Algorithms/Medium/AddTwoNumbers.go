package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		d = ListNode{2, nil}
		c = ListNode{5, &d}
		b = ListNode{2, &c}
		a = ListNode{6, &b}

		k = ListNode{3, nil}
		f = ListNode{8, &k}
		g = ListNode{1, &f}
		h = ListNode{7, &g}
	)
	current := addTwoNumbers(&a, &h)
	for current != nil {
		fmt.Print(current.Val)
		current = current.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l3                    ListNode
		current               = &l3
		val1, val2, sum, tens int
	)

	for {
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum = val1 + val2 + tens
		tens = sum / 10
		current.Val = sum % 10

		if l1 == nil && l2 == nil {
			if tens != 0 {
				current.Next = &ListNode{tens, nil}
			}
			break
		}
		current.Next = &ListNode{}
		current = current.Next
		val1 = 0
        	val2 = 0
	}

	return &l3
}
