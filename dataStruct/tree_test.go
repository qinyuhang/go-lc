package dataStruct

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMakeTreeFromLevelEmtpy(t *testing.T) {
    data := []interface{}{}
    tree := MakeTreeFromLevel(data)
    var expected *TreeNode = nil
    assert.Equal(t, tree, expected)
}

func TestMakeTreeFromLevel(t *testing.T) {
    data := []interface{}{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    //fmt.Println(*tree)
    assert.NotEqual(t, nil, tree)
    assert.Equal(t, 1, tree.Val)
    assert.Equal(t, 2, tree.Left.Val)
    assert.Equal(t, 3, tree.Right.Val)
    //fmt.Println(*tree.Left, *tree.Right)
    assert.Equal(t, 4, tree.Left.Left.Val)
    assert.Equal(t, 5, tree.Left.Right.Val)
    assert.Equal(t, 6, tree.Right.Left.Val)
    assert.Equal(t, 7, tree.Right.Right.Val)
}

func TestPreorder(t *testing.T) {
    data := []interface{}{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Preorder(tree)
    assert.Equal(t, []interface{}{1, 2, 4, 5, 3, 6, 7}, result)
}
func TestInorder(t *testing.T) {
    data := []interface{}{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Inorder(tree)
    assert.Equal(t, []interface{}{4, 2, 5, 1, 6, 3, 7}, result)
}
func TestPostorder(t *testing.T) {
    data := []interface{}{1, 2, 3, 4, 5, 6, 7}
    tree := MakeTreeFromLevel(data)
    result := Postorder(tree)
    assert.Equal(t, []interface{}{4, 5, 2, 6, 7, 3, 1}, result)
}
