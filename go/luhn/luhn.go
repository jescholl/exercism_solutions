package luhn

import "strings"

func Valid(num string) bool {
	sum := 0

	num = strings.Replace(num, " ", "", -1)
	if len(num) <= 1 {
		return false
	}

	for i := 0; i < len(num); i++ {
		position := len(num) - 1 - i
		digit := int(num[position] - '0')

		if digit < 0 || digit > 9 {
			return false
		}

		if i%2 != 0 {
			if digit = digit * 2; digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}
	return sum%10 == 0
}
