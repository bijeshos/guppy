package arrayutil_test

import (
	. "github.com/bijeshos/guppy/arrayutil"
	"testing"
)

func TestIsPresent(t *testing.T) {
	cases := []struct {
		input    []string
		search   string
		expected bool
	}{
		{[]string{"apple", "mango", "orange"}, "orange", true},
	}
	for _, c := range cases {

		actual := IsPresent(c.input, c.search)
		if actual != c.expected {
			t.Errorf("IsPresent(%v, %v) == %v, expected %v", c.input, c.search, actual, c.expected)
		}
	}
}
