package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

func TestIsSameTree(t *testing.T) {
    t1 := dataStruct.MakeTreeFromLevel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t2 := dataStruct.MakeTreeFromLevel([]int{1, 2, 3, 4, 5, 6, 7, 8})
    t3 := dataStruct.MakeTreeFromLevel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t4 := dataStruct.MakeTreeFromLevel([]int{1, 2, 3, 4, 5, 6, 7, 9})

    assert.Equal(t, IsSameTree(t1, t3), true)
    assert.Equal(t, IsSameTree(t1, t2), false)
    assert.Equal(t, IsSameTree(t1, t4), false)
    assert.Equal(t, IsSameTree(t2, t4), false)
    assert.Equal(t, IsSameTree(t3, t4), false)
}
