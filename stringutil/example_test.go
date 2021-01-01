package stringutil_test

import (
	. "github.com/bijeshos/guppy/stringutil"
	"log"
)

func ExampleRemoveSpecialChars() {
	input := "abc#$d"
	result := RemoveSpecialChars(input)
	log.Print(result)

	//Output:
	//abcd
}
