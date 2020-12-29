package arrayutil

import "testing"

func TestIsPresent(t *testing.T) {
	cases := []struct {
		in     []string
		search string
		want   bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
	}
	for _, c := range cases {
		got := IsPresent(c.in, c.search)
		if got != c.want {
			t.Errorf("IsPresent(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
