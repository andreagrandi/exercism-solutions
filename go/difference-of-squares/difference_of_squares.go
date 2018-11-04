package diffsquares

import (
	"math"
)

// SumOfSquares return the sum of squares of a given number range
func SumOfSquares(n int) int {
	var sumOfSquaresTotal int

	for i := 1; i <= n; i++ {
		sumOfSquaresTotal += int(math.Pow(float64(i), 2))
	}

	return sumOfSquaresTotal
}

// SquareOfSum return the square of sum of a given number range
func SquareOfSum(n int) int {
	var squareOfSumTotal int

	for i := 1; i <= n; i++ {
		squareOfSumTotal += i
	}

	return int(math.Pow(float64(squareOfSumTotal), 2))
}

// Difference return the difference between the SquareOfSum and the SumOfSquares of a given number range
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
