package grains

import "math"
import "errors"

func Square(square int) (grains uint64, err error) {
	if square <= 0 || square > 64 {
		err = errors.New("square out of range")
	}
	grains = uint64(math.Pow(2, float64(square-1)))
	return
}
func Total() (grains uint64) {
	for i := 1; i <= 64; i++ {
		square, err := Square(i)
		if err != nil {
			return 0
		}
		grains += square
	}
	return
}
