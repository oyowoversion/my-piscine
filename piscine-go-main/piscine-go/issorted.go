package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			ascending = false
		} else if comp < 0 {
			descending = false
		}
	}

	return ascending || descending
}
