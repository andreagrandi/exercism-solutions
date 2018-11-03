package isogram

import "unicode"

// IsIsogram checks if the input string is an isogram
func IsIsogram(input string) bool {
	m := make(map[rune]bool)

	for _, char := range input {
		char = unicode.ToLower(char)

		if unicode.IsLetter(char) {
			_, exist := m[char]

			if exist {
				return false
			}

			m[char] = true
		}
	}

	return true
}
