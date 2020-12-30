package arrayutil

import (
	"log"
	"testing"
)

func TestIsPresent(t *testing.T) {
	cases := []struct {
		in     []string
		search string
		want   bool
	}{
		{[]string{"a", "b", "c"}, "c", true},
	}
	for _, c := range cases {
		got := IsPresent(c.in, c.search)
		if got != c.want {
			t.Errorf("IsPresent(%v, %v) == %v, want %t", c.search, c.in, got, c.want)
		}
	}
}

func ExampleIsPresent() {
	arr := []string{"a", "b", "c"}
	input := "c"
	got := IsPresent(arr, input)
	log.Print("is present:", got)

}
