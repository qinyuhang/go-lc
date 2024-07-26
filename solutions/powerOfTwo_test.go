package solutions

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPowerOfTwo(t *testing.T) {
    cases := []TestCase[int, bool]{
        {1, true},
        {16, true},
        {3, false},
        {1 << 30, true},
    }

    fmt.Println(1 << 1)
    for _, v := range cases {
        assert.Equal(t, v.expected, isPowerOfTwo(v.input))
    }
}
