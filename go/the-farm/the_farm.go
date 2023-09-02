package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amt, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	ff, err2 := fc.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}

	return amt * ff / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows == 0:
		return errors.New(fmt.Sprintf("%d cows are invalid: no cows don't need food", cows))
	case cows < 0:
		return errors.New(fmt.Sprintf("%d cows are invalid: there are no negative cows", cows))
	}

	return nil
}

// TODO: define the 'DivideFood' function

// TODO: define the 'ValidateInputAndDivideFood' function

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
