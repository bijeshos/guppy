package stringutil

import "regexp"

const (
	SpecialCharsRegEx = "[~=?!+,\\$#&@'\\(\\)\\[\\]\"]"
	NonEnglishRegEx   = "\\W"
	EmptyString       = ""
)

//RemoveSpecialChars removes special characters from the input string
func RemoveSpecialChars(input string) string {
	return RemoveByPattern(input, SpecialCharsRegEx)
}

//RemoveNonEnglishChars removes non-english characters from the input string
func RemoveNonEnglishChars(input string) string {
	return RemoveByPattern(input, NonEnglishRegEx)
}

//ReplaceByPattern replaces parts of the string based on given regex pattern
func ReplaceByPattern(input, pattern, replacement string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(input, replacement)
}

//RemoveByPattern removes parts of the string based on given regex pattern
func RemoveByPattern(input, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(input, EmptyString)
}
