package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = l.Head // The tail becomes the old head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev // The head becomes the last node we processed
}
