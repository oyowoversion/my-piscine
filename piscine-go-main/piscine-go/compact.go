package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	nonEmpty := make([]string, 0, len(original))

	for _, s := range original {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}

	*ptr = nonEmpty
	return len(nonEmpty)
}
