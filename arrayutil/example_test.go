package arrayutil_test

import (
	"fmt"
	. "github.com/bijeshos/guppy/arrayutil"
)

func ExampleIsPresent() {
	input := []string{"apple", "mango", "orange"}
	search := "orange"
	result := IsPresent(input, search)
	fmt.Println(result)

	//Output:
	//true

}
