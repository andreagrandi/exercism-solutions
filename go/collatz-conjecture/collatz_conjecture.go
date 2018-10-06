package collatzconjecture

import "errors"

// CollatzConjecture divides by 2 if even, multiplies by 3 and add 1 of odd
func CollatzConjecture(input int) (int, error) {
	steps := 0

	if input < 1 {
		return steps, errors.New("Invalid number")
	}

	for input > 1 {
		// if number is even
		if input%2 == 0 {
			input = input / 2
		} else {
			input = (input * 3) + 1
		}
		steps++
	}

	return steps, nil
}
