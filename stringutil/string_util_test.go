package stringutil_test

import (
	. "github.com/bijeshos/guppy/stringutil"
	"testing"
)

func TestRemoveSpecialChars(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"abc#$d", "abcd"},
	}
	for _, c := range cases {
		actual := RemoveSpecialChars(c.input)
		if actual != c.expected {
			t.Errorf("RemoveSpecialChars(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
