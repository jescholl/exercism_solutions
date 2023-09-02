// Package leap calculates wether or not a year is a leap year
package leap

// IsLeapYear returns bool indicating wether or not a year is a leap year
func IsLeapYear(year int) bool {
	divisibleBy4 := year%4 == 0
	divisibleBy100 := year%100 == 0
	divisibleBy400 := year%400 == 0

	return divisibleBy4 && (!divisibleBy100 || divisibleBy400)
}
