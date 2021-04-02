package etl

import "strings"

type OldFormat map[int][]string
type NewFormat map[string]int

func Transform(input OldFormat) NewFormat {
	var output = NewFormat{}

	for k, v := range input {
		for _, strValue := range v {
			newKey := strings.ToLower(strValue)
			output[newKey] = k
		}
	}

	return output
}
