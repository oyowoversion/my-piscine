package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb == 3 {
		return 3
	}

	// Start with the next odd number if even
	if nb%2 == 0 {
		nb++
	}

	for {
		if isPrime(nb) {
			return nb
		}
		// Skip even numbers by adding 2
		nb += 2
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisors of form 6k ± 1 up to sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
