package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Calculate total length for performance (optional but good practice)
	totalLength := 0
	for _, s := range args {
		totalLength += len(s)
	}
	totalLength += len(args) - 1 // for '\n' between elements

	// Create a byte slice to build the string
	result := make([]byte, 0, totalLength)

	for i, word := range args {
		for j := 0; j < len(word); j++ {
			result = append(result, word[j])
		}
		if i != len(args)-1 {
			result = append(result, '\n')
		}
	}

	return string(result)
}
