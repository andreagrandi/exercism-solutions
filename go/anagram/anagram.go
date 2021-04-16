package anagram

import (
	"reflect"
	"strconv"
	"strings"
)

func getWordMap(input string) map[string]int {
	inputMap := make(map[string]int)

	for _, c := range input {
		cStr := strconv.QuoteRune(c)
		inputMap[cStr] += 1
	}

	return inputMap
}
func Detect(input string, words []string) []string {
	input = strings.ToLower(input)
	inputMap := getWordMap(input)
	anagrams := []string{}

	for _, word := range words {
		checkWord := strings.ToLower(word)

		// anagrams of themselves are not valid
		if input != checkWord {
			wordMap := getWordMap(checkWord)
			if reflect.DeepEqual(inputMap, wordMap) {
				anagrams = append(anagrams, word)
			}
		}
	}

	return anagrams
}
