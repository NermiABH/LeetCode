/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }else if head.Next.Next == nil {
        head.Next = nil
        return head
    }

    l1, l2 := head, head.Next.Next

    for l2.Next != nil && l2.Next.Next != nil {
        l1 = l1.Next
        l2 = l2.Next.Next
    }
    if l2.Next == nil {
        l1.Next = l1.Next.Next
    }else {
        l1.Next.Next = l1.Next.Next.Next     
    }

    return head
}