/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    l1, l2 := head, head.Next
    headEven := l2

    for l1.Next != nil && l2.Next != nil {
        if l1.Next != nil && l1.Next.Next != nil {
            l1.Next = l1.Next.Next
            l1 = l1.Next
        }else {
            l2.Next = nil
        }

        if l2.Next != nil && l2.Next.Next != nil {
            l2.Next = l2.Next.Next
            l2 = l2.Next
        }else {
            l2.Next = nil
        }
    }
    l1.Next = headEven
    return head
}