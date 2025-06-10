package piscine

func JumpOver(str string) string {
	result := ""
	// Iterate through the string and collect every third character
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// If the result is empty, return a newline
	if result == "" {
		return "\n"
	}

	return result + "\n"
}
