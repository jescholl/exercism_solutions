package isogram

import "strings"

func IsIsogram(phrase string) bool {
	letterCount := map[byte]int{}
	isIsogram := true
	phrase = strings.ToUpper(phrase)

	for i := 0; i < len(phrase); i++ {
		letter := phrase[i]

		if letter >= 'A' && letter <= 'Z' {
			letterCount[letter]++
		}
		if letterCount[letter] > 1 {
			isIsogram = false
		}
	}

	return isIsogram
}
