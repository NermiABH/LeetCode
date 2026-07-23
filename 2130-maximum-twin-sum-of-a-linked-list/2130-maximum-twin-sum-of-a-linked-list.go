/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    l, r := head.Next, head.Next
    for r.Next != nil {
        l, r = l.Next, r.Next.Next
    }

    var prev *ListNode
    curr := l
    for curr != nil {
        next := curr.Next
        curr.Next = prev

        prev = curr
        curr = next
    }


    l, r = head, prev
    var maxTwin int
    for r != nil {
        maxTwin = max(maxTwin, l.Val+r.Val)
        l, r = l.Next, r.Next
    }
    
    return maxTwin
}