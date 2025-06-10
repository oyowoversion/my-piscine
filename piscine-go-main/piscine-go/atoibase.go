package piscine

func AtoiBase(s string, base string) int {
	if !isValidBaseForAtoiBase(base) {
		return 0
	}

	baseRunes := []rune(base)
	baseLen := len(baseRunes)
	valueMap := make(map[rune]int)

	for i, r := range baseRunes {
		valueMap[r] = i
	}

	result := 0
	for _, r := range s {
		val, exists := valueMap[r]
		if !exists {
			return 0
		}
		result = result*baseLen + val
	}

	return result
}

func isValidBaseForAtoiBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}
