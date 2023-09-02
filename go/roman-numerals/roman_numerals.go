package romannumerals

import (
	"errors"
)

type Numeral map[int]string

var numerals = map[int]Numeral{
	1: {
		1:  "I",
		5:  "V",
		10: "X",
	},
	10: {
		1:  "X",
		5:  "L",
		10: "C",
	},
	100: {
		1:  "C",
		5:  "D",
		10: "M",
	},
	1000: {
		1:  "M",
		5:  "",
		10: "",
	},
}

func ToRomanNumeral(input int) (string, error) {
	var output string

	if input > 3999 || input < 1 {
		return "", errors.New("Invalid input")
	}

	for d := 1000; d > 0; d /= 10 {
		digit := int(input / d)
		m := numerals[d]

		input -= digit * d

		switch digit {
		case 0:
			// no zero's in Roman Numerals
		case 1:
			output += m[1]
		case 2:
			output += m[1] + m[1]
		case 3:
			output += m[1] + m[1] + m[1]
		case 4:
			output += m[1] + m[5]
		case 5:
			output += m[5]
		case 6:
			output += m[5] + m[1]
		case 7:
			output += m[5] + m[1] + m[1]
		case 8:
			output += m[5] + m[1] + m[1] + m[1]
		case 9:
			output += m[1] + m[10]
		default:
			return "", errors.New("Internal logic error")
		}
	}

	return output, nil
}
