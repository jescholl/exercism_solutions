package pangram

import "strings"

func IsPangram(input string) bool {
	lettersSeen := make([]bool, 'z'+1)

	for _, chr := range strings.ToLower(input) {
		lettersSeen[chr] = true
	}

	for _, seen := range lettersSeen['a':'z'] {
		if !seen {
			return false
		}
	}
	return true
}
