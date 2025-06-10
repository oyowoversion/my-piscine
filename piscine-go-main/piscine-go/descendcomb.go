package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

// helper to print two-digit numbers using PrintRune
func printTwoDigits(n int) {
	tens := n / 10
	ones := n % 10
	z01.PrintRune(rune(tens + '0'))
	z01.PrintRune(rune(ones + '0'))
}
