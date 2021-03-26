package reverse

func Reverse(s string) string {
	output := ""

	for _, c := range s {
		output = string(c) + output
	}

	return output
}
