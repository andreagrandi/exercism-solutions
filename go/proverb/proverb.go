// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := make([]string, 0)

	for i, word := range rhyme {
		if i < len(rhyme)-1 {
			sentence := fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[i+1])
			proverb = append(proverb, sentence)
		} else {
			sentence := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			proverb = append(proverb, sentence)
		}
	}

	return proverb
}
