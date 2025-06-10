package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{
		Data: data,
		Next: l.Head, // New node points to current head
	}

	l.Head = newNode // Update head to new node

	// If list was empty, also update tail
	if l.Tail == nil {
		l.Tail = newNode
	}
}
