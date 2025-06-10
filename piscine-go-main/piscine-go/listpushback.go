package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		// List is empty; new node is both Head and Tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// List is not empty; append to Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
