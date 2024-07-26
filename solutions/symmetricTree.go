package solutions

func isMirror(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
func IsSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil || root.Right == nil {
        return false
    }
    return isMirror(root.Left, root.Right)
}
