package dataStruct

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMakeTreeFromLevelEmtpy(t *testing.T) {
    data := []int{}
    tree := MakeTreeFromLevel(data)
    var expected *TreeNode = nil
    assert.Equal(t, tree, expected)
}

func TestMakeTreeFromLevel(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    //fmt.Println(*tree)
    assert.NotEqual(t, tree, nil)
    assert.Equal(t, tree.Val, 1)
    assert.Equal(t, tree.Left.Val, 2)
    assert.Equal(t, tree.Right.Val, 3)
    //fmt.Println(*tree.Left, *tree.Right)
    assert.Equal(t, tree.Left.Left.Val, 4)
    assert.Equal(t, tree.Left.Right.Val, 5)
    assert.Equal(t, tree.Right.Left.Val, 6)
    assert.Equal(t, tree.Right.Right.Val, 7)
}

func TestPreorder(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Preorder(tree)
    fmt.Println(result)
    assert.Equal(t, result, []int{1, 2, 4, 5, 3, 6, 7})
}
func TestInorder(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Inorder(tree)
    fmt.Println(result)
    assert.Equal(t, result, []int{4, 2, 5, 1, 6, 3, 7})
}
func TestPostorder(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Postorder(tree)
    fmt.Println(result)
    assert.Equal(t, result, []int{4, 5, 2, 6, 7, 3, 1})
}
