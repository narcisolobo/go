package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float64
	switch {
	case balance < 0:
		interestRate = 3.213
	case balance >= 0 && balance < 1000:
		interestRate = .5
	case balance >= 1000 && balance < 5000:
		interestRate = 1.621
	case balance >= 5000:
		interestRate = 2.475
	}
	return float32(interestRate)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		years++
		balance = AnnualBalanceUpdate(balance)
	}
	return years
}
