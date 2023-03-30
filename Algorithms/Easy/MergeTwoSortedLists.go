package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := &ListNode{}
	curList3 := list3
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curList3.Next = list1
			list1 = list1.Next
		} else {
			curList3.Next = list2
			list2 = list2.Next
		}
		curList3 = curList3.Next
	}
	if list1 == nil {
		curList3.Next = list2
	} else if list2 == nil {
		curList3.Next = list1
	}
	return list3.Next
}

func mergeTwoListsRec(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
