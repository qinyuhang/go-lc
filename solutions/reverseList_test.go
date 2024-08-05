package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

func TestReverseList(t *testing.T) {
    cases := []TestCase[[]int, []int]{
        {[]int{1, 2, 3}, []int{3, 2, 1}},
        {[]int{}, []int{}},
    }
    for _, v := range cases {
        list := dataStruct.SliceToListNode(v.input)
        result := reverseList(list)
        resultArr := dataStruct.ListNodeToSlice(result)
        assert.Equal(t, v.expected, resultArr)
    }
}
