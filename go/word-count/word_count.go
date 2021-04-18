package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	frequency := make(Frequency)
	regex := regexp.MustCompile(`\w+'?\w+|\w`)
	words := regex.FindAllString(strings.ToLower(input), -1)

	for _, word := range words {
		frequency[word] += 1
	}

	return frequency
}
