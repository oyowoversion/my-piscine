package piscine

func StringToIntSlice(str string) []int {
	var result []int

	// Convert the string to a slice of runes (handles multi-byte characters)
	for _, char := range str {
		result = append(result, int(char)) // Convert each rune to its Unicode code point
	}

	return result
}
