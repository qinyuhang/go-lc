package dataStruct

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestListNodeToSlice(t *testing.T) {
    cases := []*TestCase[[]int, []int]{
        {[]int{1, 2, 3}, []int{1, 2, 3}},
        {[]int{}, []int{}},
    }
    for _, v := range cases {
        list := SliceToListNode(v.input)
        out := ListNodeToSlice(list)
        assert.Equal(t, v.expected, out)
    }
}
