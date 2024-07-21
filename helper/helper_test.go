package helper

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSlicesEqual(t *testing.T) {
    a := []int{1, 2}
    b := []int{1, 2}
    c := []int{2, 1}
    d := []int{}

    assert.Equal(t, true, SlicesEqual(a, b))
    assert.Equal(t, false, SlicesEqual(a, c))
    assert.Equal(t, false, SlicesEqual(b, c))
    assert.Equal(t, false, SlicesEqual(d, c))
}
