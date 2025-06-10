package piscine

func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			// rotate lowercase letters
			r = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			// rotate uppercase letters
			r = 'A' + (r-'A'+14)%26
		}
		result = append(result, r)
	}

	return string(result)
}
