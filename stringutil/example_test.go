package stringutil_test

import (
	"fmt"
	. "github.com/bijeshos/guppy/stringutil"
)

func ExampleRemoveSpecialChars() {
	input := "abc#$d"
	result := RemoveSpecialChars(input)
	fmt.Println(result)

	//Output:
	//abcd
}
