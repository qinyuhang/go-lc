package dataStruct

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func MakeTreeFromLevel(data []int) *TreeNode {
    if len(data) == 0 {
        return nil
    }

    root := &TreeNode{Val: data[0]}
    queue := []*TreeNode{}

    dataWithoutFirst := data[1:]
    cur := root

    for _, v := range dataWithoutFirst {
        tmp := &TreeNode{Val: v}
        queue = append(queue, tmp)
        if cur.Left == nil {
            cur.Left = tmp
        } else if cur.Right == nil {
            cur.Right = tmp
            cur = queue[0]
            queue = queue[1:]
        }
    }
    return root
}

func Preorder(t *TreeNode) []int {
    result := make(chan int)
    go func() {
        preorderHelper(t, result)
        close(result)
    }()
    var values []int
    for val := range result {
        values = append(values, val)
    }
    return values
}

func preorderHelper(t *TreeNode, result chan int) {
    if t == nil {
        return
    }
    result <- t.Val
    preorderHelper(t.Left, result)
    preorderHelper(t.Right, result)
}

func Inorder(t *TreeNode) []int {
    result := make(chan int)
    go func() {
        inorderHelper(t, result)
        close(result)
    }()
    var values []int
    for val := range result {
        values = append(values, val)
    }
    return values
}

func inorderHelper(t *TreeNode, result chan int) {
    if t == nil {
        return
    }
    inorderHelper(t.Left, result)
    result <- t.Val
    inorderHelper(t.Right, result)
}

func Postorder(t *TreeNode) []int {
    result := make(chan int)
    go func() {
        postorderHelper(t, result)
        close(result)
    }()
    var values []int
    for val := range result {
        values = append(values, val)
    }
    return values
}

func postorderHelper(t *TreeNode, result chan int) {
    if t == nil {
        return
    }
    postorderHelper(t.Left, result)
    postorderHelper(t.Right, result)
    result <- t.Val
}
