// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// IsTextUpperCase checks if the whole string is written in upper case
func IsTextUpperCase(input string) bool {
	isUpper := false

	for _, char := range input {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				isUpper = true
			} else {
				isUpper = false
				return false
			}
		}
	}

	return isUpper
}

// IsTextAQuestion checks if the text is a question ('?' must be at the end, excluding spaces)
func IsTextAQuestion(input string) bool {
	isQuestion := false
	input = strings.TrimSpace(input)

	if len(input) > 0 {
		lastChar := input[len(input)-1:]

		if lastChar == "?" {
			isQuestion = true
		}
	}

	return isQuestion
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	} else if IsTextAQuestion(remark) && IsTextUpperCase(remark) {
		return "Calm down, I know what I'm doing!"
	} else if IsTextAQuestion(remark) && !IsTextUpperCase(remark) {
		return "Sure."
	} else if !IsTextAQuestion(remark) && IsTextUpperCase(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
