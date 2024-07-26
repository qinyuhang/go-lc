package solutions

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    l := root.Left
    root.Left = root.Right
    root.Right = l
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

func InvertTree(root *TreeNode) *TreeNode {
    return invertTree(root)
}
