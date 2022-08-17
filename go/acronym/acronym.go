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
	acronym := ""

	// sanitise "-"
	s = strings.ReplaceAll(s, "-", " ")

	// remove all special chars
	onlyLettersString := ""
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsSpace(c) {
			onlyLettersString += string(c)
		}
	}

	// split the string into words
	words := strings.Split(onlyLettersString, " ")

	// for each word we remove spaces, uppercase it and take the first letter
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 0 {
			word = strings.ToUpper(word)
			acronym += string(word[0])
		}
	}

	return acronym
}
