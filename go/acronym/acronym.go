// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	lowerInput := strings.ToLower(s)
	titledInput := strings.Title(lowerInput)

	acronym := ""

	for _, char := range titledInput {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				acronym = acronym + string(char)
			}
		}
	}

	return acronym
}
