package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Special case for 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract digits
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Sort digits (simple bubble sort)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Print sorted digits
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
