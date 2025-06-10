package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := false
	foundDigit := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			foundDigit = true
			result = result*10 + int(char-'0')
		} else if char == '-' && !foundDigit {
			negative = true
		}
	}

	if negative {
		return -result
	}
	return result
}
