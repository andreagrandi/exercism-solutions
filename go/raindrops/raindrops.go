package raindrops

import (
	"strconv"
)

// Convert accepts an integer and return a PlingPlangPlong string
func Convert(input int) string {
	var converted string

	if input%3 == 0 {
		converted += "Pling"
	}

	if input%5 == 0 {
		converted += "Plang"
	}

	if input%7 == 0 {
		converted += "Plong"
	}

	if converted == "" {
		converted = strconv.Itoa(input)
	}

	return converted
}
