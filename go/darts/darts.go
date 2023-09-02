package darts

import "math"

func Score(x, y float64) int {
	score := 0

	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case distance <= 1:
		score = 10
	case distance <= 5:
		score = 5
	case distance <= 10:
		score = 1
	}

	return score
}
