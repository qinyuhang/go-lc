package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

func TestIsSymmetric(t *testing.T) {
    case1 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 2, 3, 4, 4, 3})
    case2 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 2, nil, 3, nil, 3})

    assert.Equal(t, true, IsSymmetric(case1))
    assert.Equal(t, false, IsSymmetric(case2))

    case3 := dataStruct.MakeTreeFromLevel([]interface{}{})
    assert.Equal(t, true, IsSymmetric(case3))

    case4 := dataStruct.MakeTreeFromLevel([]interface{}{1, nil, 2})
    assert.Equal(t, false, IsSymmetric(case4))
    case5 := dataStruct.MakeTreeFromLevel([]interface{}{1})
    assert.Equal(t, true, IsSymmetric(case5))

    case6 := dataStruct.MakeTreeFromLevel([]interface{}{1, 2, 3})
    assert.Equal(t, false, IsSymmetric(case6))
}
