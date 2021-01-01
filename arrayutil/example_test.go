package arrayutil_test

import (
	. "github.com/bijeshos/guppy/arrayutil"
	"log"
)

func ExampleIsPresent() {
	input := []string{"apple", "mango", "orange"}
	search := "orange"
	result := IsPresent(input, search)
	log.Print("is present:", result)

}
