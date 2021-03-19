package romannumerals

import (
	"errors"
	"strings"
)

type Hundreds map[int]string
type Tens map[int]string
type Units map[int]string

var hundreds = Hundreds{
	0: "",
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}

var tens = Tens{
	0: "",
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

var units = Units{
	0: "",
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

func ToRomanNumeral(number int) (string, error) {
	if number > 3000 {
		return "", errors.New("number is too big: max number allowed is 3000")
	}

	if number <= 0 {
		return "", errors.New("number must be greater than 0")
	}

	output := ""

	// Handle thousands
	numberOfThousands := number / 1000
	output += strings.Repeat("M", numberOfThousands)
	number = number - (numberOfThousands * 1000)

	// Handle hundreds
	numberOfHundreds := number / 100
	output += hundreds[numberOfHundreds]
	number = number - (numberOfHundreds * 100)

	// Handle tens
	numberOfTens := number / 10
	output += tens[numberOfTens]
	number = number - (numberOfTens * 10)

	// Handle units
	output += units[number]

	return output, nil
}
