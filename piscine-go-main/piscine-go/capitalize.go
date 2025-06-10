package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	capNext := true

	for i, char := range runes {
		if isAlphaNumeric(char) {
			if capNext {
				if char >= 'a' && char <= 'z' {
					runes[i] = char - 32 // to uppercase
				}
				capNext = false
			} else {
				if char >= 'A' && char <= 'Z' {
					runes[i] = char + 32 // to lowercase
				}
			}
		} else {
			capNext = true
		}
	}
	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') ||
		(r >= 'a' && r <= 'z') ||
		(r >= '0' && r <= '9')
}
