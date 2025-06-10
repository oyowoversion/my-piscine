package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := -1

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				words = append(words, s[start:i])
				start = -1
			}
		}
	}

	// Add the last word if the string doesn't end with whitespace
	if start != -1 {
		words = append(words, s[start:])
	}

	return words
}
