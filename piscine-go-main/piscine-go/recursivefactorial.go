package piscine

func RecursiveFactorial(nb int) int {
	// Handle invalid cases (negative numbers or potential overflow)
	if nb < 0 || nb > 20 {
		return 0
	}
	// Base case: 0! and 1! are both 1
	if nb == 0 || nb == 1 {
		return 1
	}
	// Recursive case: nb! = nb * (nb-1)!
	result := nb * RecursiveFactorial(nb-1)
	// Check for overflow
	if result < 0 {
		return 0
	}
	return result
}
