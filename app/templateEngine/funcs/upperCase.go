package templatefuncs

import "strings"

// CustomUppercase is a custom template function that converts a given string to uppercase.
//
// Example Usage:
//
//	{{ customUppercase "hello" }} // Outputs: "HELLO"
func CustomUppercase(str string) string {
	return strings.ToUpper(str)
}
