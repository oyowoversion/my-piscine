package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Create a slice by copying the input
	corrected := podium[:]

	// Reverse the order manually
	for i := 0; i < len(corrected)/2; i++ {
		corrected[i], corrected[len(corrected)-1-i] = corrected[len(corrected)-1-i], corrected[i]
	}

	return corrected
}
