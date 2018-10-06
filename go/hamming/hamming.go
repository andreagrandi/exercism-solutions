package hamming

import "errors"

// Distance calculate how many differences there are between two strings of the same lenght
func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return -1, errors.New("Strings don't have the same lenght")
	}

	for i := 0; i <= len(a)-1; i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
