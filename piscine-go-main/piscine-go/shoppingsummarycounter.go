package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else {
			result[word]++
			word = ""
		}
	}
	// Add the last word (even if empty)
	result[word]++
	return result
}
