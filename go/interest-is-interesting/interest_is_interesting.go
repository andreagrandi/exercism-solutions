package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default: // when balance is >= 5000
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance / 100.0 * float64(InterestRate(balance))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// In case the balance is already at the target, no years need to pass to reach it
	if balance >= targetBalance {
		return 0
	}

	for years := 1; ; years++ {
		balance = balance + Interest(balance)
		if balance >= targetBalance {
			return years
		}
	}
}
