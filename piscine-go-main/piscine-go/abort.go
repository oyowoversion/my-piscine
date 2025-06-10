package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	// Simple bubble sort (since we have only 5 elements)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr[2] // Return the median (middle value)
}
