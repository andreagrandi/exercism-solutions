package pangram

import "strings"

func IsPangram(input string) bool {
	var alphabet = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"i": false,
		"j": false,
		"k": false,
		"l": false,
		"m": false,
		"n": false,
		"o": false,
		"p": false,
		"q": false,
		"r": false,
		"s": false,
		"t": false,
		"u": false,
		"v": false,
		"w": false,
		"x": false,
		"y": false,
		"z": false,
	}

	counter := 0

	for _, char := range input {
		letter := strings.ToLower(string(char))
		if value, ok := alphabet[letter]; ok {
			if !value {
				alphabet[letter] = true
				counter += 1
			}
		}
	}

	return counter == len(alphabet)
}
