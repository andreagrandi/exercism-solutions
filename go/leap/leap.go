package leap

// IsLeapYear check if the year is leap or not
func IsLeapYear(year int) bool {
	leap := false

	if year%4 == 0 {
		leap = true

		if year%100 == 0 {
			leap = false

			if year%400 == 0 {
				leap = true
			}
		}
	}

	return leap
}
