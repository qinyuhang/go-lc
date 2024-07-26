package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

func TestIsSameTree(t *testing.T) {

    t1 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t2 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 3, 4, 5, 6, 7, 8})
    t3 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t4 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 3, 4, 5, 6, 7, 9})

    cases := []TestCase[
        []*dataStruct.TreeNode,
        bool,
    ]{
        {
            input: []*dataStruct.TreeNode{t1, t3}, expected: true,
        }, {
            input: []*dataStruct.TreeNode{t1, t2}, expected: false,
        }, {
            input: []*dataStruct.TreeNode{t1, t4}, expected: false,
        }, {
            input: []*dataStruct.TreeNode{t2, t4}, expected: false,
        }, {
            input: []*dataStruct.TreeNode{t3, t4}, expected: false,
        },
    }

    for _, v := range cases {
        assert.Equal(t, v.expected, IsSameTree(v.input[0], v.input[1]))
    }
    assert.Equal(t, IsSameTree(t1, t3), true)
    assert.Equal(t, IsSameTree(t1, t2), false)
    assert.Equal(t, IsSameTree(t1, t4), false)
    assert.Equal(t, IsSameTree(t2, t4), false)
    assert.Equal(t, IsSameTree(t3, t4), false)
}
