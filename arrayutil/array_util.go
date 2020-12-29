package arrayutil

import "fmt"

func IsPresent(array []string, value string) bool {
	fmt.Println(array, value)
	_, ok := Find(array, value)
	return ok
}

func Find(array []string, value string) (int, bool) {
	for i, item := range array {
		if item == value {
			return i, true
		}
	}
	return -1, false
}
