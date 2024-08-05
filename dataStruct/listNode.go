package dataStruct

type ListNode struct {
    Val  int
    Next *ListNode
}

func SliceToListNode(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    fakeHead := &ListNode{Val: -1}
    cur := fakeHead
    for _, v := range arr {
        tmp := &ListNode{Val: v}
        cur.Next = tmp
        cur = cur.Next
    }
    return fakeHead.Next
}

func ListNodeToSlice(head *ListNode) []int {
    result := []int{}
    if head == nil {
        return result
    }
    cur := head
    for cur != nil {
        result = append(result, cur.Val)
        cur = cur.Next
    }

    return result
}
