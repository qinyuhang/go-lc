package solutions

import "go_lc/dataStruct"

type ListNode = dataStruct.ListNode

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var prev *ListNode = nil
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}
