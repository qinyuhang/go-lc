package dataStruct

type TestCase[T any, P any] struct {
    input    T
    expected P
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func MakeTreeFromLevel(data []interface{}) *TreeNode {
    if len(data) == 0 || data[0] == nil {
        return nil
    }
    root := &TreeNode{Val: data[0].(int)}
    queue := []*TreeNode{root}
    index := 1
    for len(queue) > 0 && index < len(data) {
        current := queue[0]
        queue = queue[1:]
        // 处理左子节点
        if index < len(data) && data[index] != nil {
            current.Left = &TreeNode{Val: data[index].(int)}
            queue = append(queue, current.Left)
        }
        index++
        // 处理右子节点
        if index < len(data) && data[index] != nil {
            current.Right = &TreeNode{Val: data[index].(int)}
            queue = append(queue, current.Right)
        }
        index++
    }
    return root
}

func Preorder(t *TreeNode) []interface{} {
    result := make(chan interface{})
    go func() {
        preorderHelper(t, result)
        close(result)
    }()
    var values []interface{}
    for val := range result {
        values = append(values, val)
    }
    return values
}

func preorderHelper(t *TreeNode, result chan interface{}) {
    if t == nil {
        return
    }
    result <- t.Val
    preorderHelper(t.Left, result)
    preorderHelper(t.Right, result)
}

func Inorder(t *TreeNode) []interface{} {
    result := make(chan interface{})
    go func() {
        inorderHelper(t, result)
        close(result)
    }()
    var values []interface{}
    for val := range result {
        values = append(values, val)
    }
    return values
}

func inorderHelper(t *TreeNode, result chan interface{}) {
    if t == nil {
        return
    }
    inorderHelper(t.Left, result)
    result <- t.Val
    inorderHelper(t.Right, result)
}

func Postorder(t *TreeNode) []interface{} {
    result := make(chan interface{})
    go func() {
        postorderHelper(t, result)
        close(result)
    }()
    var values []interface{}
    for val := range result {
        values = append(values, val)
    }
    return values
}

func postorderHelper(t *TreeNode, result chan interface{}) {
    if t == nil {
        return
    }
    postorderHelper(t.Left, result)
    postorderHelper(t.Right, result)
    result <- t.Val
}
