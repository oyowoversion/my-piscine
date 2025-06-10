package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert from baseFrom to decimal
	num := 0
	baseLen := len(baseFrom)

	// Map baseFrom chars to values
	fromMap := make(map[rune]int)
	for i, ch := range baseFrom {
		fromMap[ch] = i
	}

	for _, ch := range nbr {
		num = num*baseLen + fromMap[ch]
	}

	// Edge case: 0
	if num == 0 {
		return string(baseTo[0])
	}

	// Convert from decimal to baseTo
	toBaseLen := len(baseTo)
	result := ""

	for num > 0 {
		remainder := num % toBaseLen
		result = string(baseTo[remainder]) + result
		num = num / toBaseLen
	}

	return result
}
