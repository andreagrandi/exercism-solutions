package accumulate

func Accumulate(sequence []string, converter func(string) string) []string {
	output := []string{}

	for _, input := range sequence {
		output = append(output, converter(input))
	}

	return output
}
