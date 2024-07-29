package solutions

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCountPrimes(t *testing.T) {
    cases := []TestCase[int, int]{
        {10, 4}, {0, 0}, {1, 0},
    }

    for _, v := range cases {
        assert.Equal(t, v.expected, countPrimes(v.input))
    }
}
