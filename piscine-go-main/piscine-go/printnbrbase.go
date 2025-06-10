package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseRunes := []rune(base)
	baseLen := len(baseRunes)

	if nbr < 0 {
		z01.PrintRune('-')
		unbr := uint(-nbr)
		printNbrBaseRec(unbr, baseRunes, baseLen)
	} else {
		printNbrBaseRec(uint(nbr), baseRunes, baseLen)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	baseMap := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || baseMap[r] {
			return false
		}
		baseMap[r] = true
	}
	return true
}

func printNbrBaseRec(n uint, base []rune, baseLen int) {
	if n >= uint(baseLen) {
		printNbrBaseRec(n/uint(baseLen), base, baseLen)
	}
	z01.PrintRune(base[n%uint(baseLen)])
}
