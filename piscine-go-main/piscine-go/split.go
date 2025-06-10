package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	start := 0
	sepLen := len(sep)

	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1 // skip the rest of sep
		}
	}
	result = append(result, s[start:]) // add the last part
	return result
}
