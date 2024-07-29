package solutions

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsomorphicString(t *testing.T) {
    cases := []TestCase[[]string, bool]{
        {[]string{"egg", "add"}, true},
        {[]string{"foo", "bar"}, false},
        {[]string{"paper", "title"}, true},
        {[]string{"a", "ab"}, false},
        {[]string{"pa", "tt"}, false},
    }

    for _, v := range cases {
        assert.Equal(t, v.expected, isIsomorphic(v.input[0], v.input[1]), "except eq", v.input)
    }
}
