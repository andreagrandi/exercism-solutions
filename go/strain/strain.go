package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(pred func(int) bool) Ints {
	if ints == nil {
		return nil
	}

	output := make(Ints, 0)

	for _, i := range ints {
		if pred(i) {
			output = append(output, i)
		}
	}

	return output
}

func (ints Ints) Discard(pred func(int) bool) Ints {
	if ints == nil {
		return nil
	}

	output := make(Ints, 0)

	for _, i := range ints {
		if !pred(i) {
			output = append(output, i)
		}
	}

	return output
}

func (lists Lists) Keep(pred func([]int) bool) Lists {
	output := make(Lists, 0)

	for _, l := range lists {
		if pred(l) {
			output = append(output, l)
		}
	}

	return output
}

func (strings Strings) Keep(pred func(string) bool) Strings {
	output := make(Strings, 0)

	for _, s := range strings {
		if pred(s) {
			output = append(output, s)
		}
	}

	return output
}
