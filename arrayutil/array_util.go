package arrayutil

//IsPresent checks whether the input string is present in the given string array
func IsPresent(arr []string, input string) bool {
	_, ok := FindIndex(arr, input)
	return ok
}

//FindIndex returns the index of input string in the given string array. Returns -1 if not present
func FindIndex(arr []string, input string) (int, bool) {
	for i, item := range arr {
		if item == input {
			return i, true
		}
	}
	return -1, false
}
