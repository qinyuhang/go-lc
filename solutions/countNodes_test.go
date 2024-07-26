package solutions

import (
    "github.com/stretchr/testify/assert"
    "go_lc/dataStruct"
    "testing"
)

func TestCountNodes(t *testing.T) {
    cases := []TestCase[[]interface{}, int]{
        {
            input:    []interface{}{1, 2, 3, 4, 5, 6},
            expected: 6,
        }, {
            input:    []interface{}{},
            expected: 0,
        }, {
            input:    []interface{}{1},
            expected: 1,
        },
    }

    for _, v := range cases {
        tr := dataStruct.MakeTreeFromLevel(v.input)
        result := countNodes(tr)
        assert.Equal(t, v.expected, result)
    }
}
