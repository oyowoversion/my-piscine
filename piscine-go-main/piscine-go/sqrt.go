package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	for i := 1; i <= nb; i++ {
		square := i * i
		if square == nb {
			return i
		}
		if square > nb {
			return 0
		}
	}
	return 0
}
