package stringutil

import (
	"log"
	"testing"
)

func TestRemoveSpecialChars(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"abc#$d", "abcd"},
	}
	for _, c := range cases {
		got := RemoveSpecialChars(c.in)
		if got != c.want {
			t.Errorf("RemoveSpecialChars(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func ExampleRemoveSpecialChars() {
	input := "abc#$d"
	got := RemoveSpecialChars(input)
	log.Print("after removing special chars:", got)

}
