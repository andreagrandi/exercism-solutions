package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cows int
}

func (sne *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sne.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// Define a custom error type called SillyNephewError.
	// It should be returned in case the number of cows is negative.
	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	// Get the amount of food and any eventual error
	amount, err := weightFodder.FodderAmount()

	// After getting the fodder amount from weightFodder,
	// prevent a division by zero when there are no cows at all by returning an error saying "division by zero"
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	// If ErrScaleMalfunction is returned by FodderAmount and the fodder amount is positive,
	// double the fodder amount returned by FodderAmount before dividing it equally between the cows.
	if amount > 0 && err == ErrScaleMalfunction {
		amount *= 2
		return amount / float64(cows), nil
	}

	// If the scale is broken and returning negative amounts of fodder,
	// return an error saying "negative fodder" only if FodderAmount returned ErrScaleMalfunction or nil
	if amount < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	}

	// For any other error besides ErrScaleMalfunction, return 0 and the error.
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}

	return amount / float64(cows), nil
}
