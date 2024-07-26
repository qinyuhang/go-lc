package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

type TestCase[T any, P any] struct {
    input    T
    expected P
}

func TestInvertTree(t *testing.T) {
    cases := []*TestCase[[]interface{}, []interface{}]{
        {
            []interface{}{4, 2, 7, 1, 3, 6, 9},
            []interface{}{4, 7, 2, 9, 6, 3, 1},
        },
        {
            []interface{}{2, 1, 3},
            []interface{}{2, 3, 1},
        },
        {
            []interface{}{},
            []interface{}{},
        },
    }
    for _, v := range cases {
        expected := dataStruct.MakeTreeFromLevel(v.expected)
        result := InvertTree(dataStruct.MakeTreeFromLevel(v.input))
        assert.Equal(t,
            dataStruct.Preorder(expected),
            dataStruct.Preorder(result),
        )
    }
}
