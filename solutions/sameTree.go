package solutions

import "go_lc/dataStruct"

type TreeNode = dataStruct.TreeNode

func IsSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
