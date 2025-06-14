package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		prev := result
		result *= i
		if result < prev {
			return 0
		}
	}
	return result
}
