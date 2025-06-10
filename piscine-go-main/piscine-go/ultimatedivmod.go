package piscine

func UltimateDivMod(a *int, b *int) {
	dividend := *a
	divisor := *b
	*a = dividend / divisor
	*b = dividend % divisor
}
