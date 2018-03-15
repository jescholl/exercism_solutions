package robotname

import (
	"math/rand"
	"time"
)

var AvailableNames []int

type Robot struct {
	name string
}

// convert number to radix as defined by the character set between minChr and maxChr
func encodeName(number, minChr, maxChr, length int) (string, int) {
	radix := maxChr - minChr + 1
	letters := make([]rune, length)
	for i := length - 1; i >= 0; i-- {
		letters[i] = rune(number%radix + minChr)
		number /= radix
	}
	// when letters is the desired length, return the remainder if there is one
	return string(letters), number
}

func (r *Robot) Name() string {
	var num int
	if r.name == "" {
		// seed rand and fill AvailableNames with randomly shuffled numbers corresponding to all possible names
		if len(AvailableNames) == 0 {
			rand.Seed(time.Now().UTC().UnixNano())
			AvailableNames = rand.Perm(26 * 26 * 10 * 10 * 10)
		}
		// pop number off list
		num, AvailableNames = AvailableNames[0], AvailableNames[1:]

		// encode the number as a 3-digit name using numbers 0-9
		numbers, remainder := encodeName(num, '0', '9', 3)
		// encode the remainder as a 2-character name using letters A-Z
		letters, remainder := encodeName(remainder, 'A', 'Z', 2)

		if remainder == 0 {
			r.name = letters + numbers
		} else {
			r.name = "Invalid name"
		}
	}

	return r.name
}

func (r *Robot) Reset() *Robot {
	r.name = ""
	return r
}
