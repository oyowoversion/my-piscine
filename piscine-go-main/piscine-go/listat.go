package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	// Handle negative positions
	if pos < 0 {
		return nil
	}

	current := l
	count := 0

	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next
	}
	return nil
}
