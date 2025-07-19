package game

type DoublyLinkedListNode struct {
	dot      *Dot
	nextNode *DoublyLinkedListNode
	prevNode *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (l *DoublyLinkedList) InsertAtEnd(dot *Dot) {
	node := &DoublyLinkedListNode{dot: dot}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.nextNode = node
		node.prevNode = l.tail
		l.tail = node
	}
}

func (l *DoublyLinkedList) RemoveByCoordinates(point Point) {
	var node *DoublyLinkedListNode
	current := l.head
	for current != nil {
		if current.dot.X == point.X && current.dot.Y == point.Y {
			node = current
			break
		}
		current = current.nextNode
	}

	if node != nil {
		if l.tail == node {
			l.tail = node.prevNode
		}
		if l.head == node {
			l.head = node.nextNode
		}
		if node.prevNode != nil {
			node.prevNode.nextNode = node.nextNode
		}
		if node.nextNode != nil {
			node.nextNode.prevNode = node.prevNode
		}

	}
}
