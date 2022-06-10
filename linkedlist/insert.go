package linkedlist

func (l *LL) InsertAtEnd(value int) {
	if l.GetHead() == nil {
		l.SetHead(NewNode(value))
		return
	}

	// go to the very end
	temp := l.GetHead()
	for temp.Next != nil {
		temp = temp.Next
	}

	// link to last node
	temp.Next = NewNode(value)
}

func (l *LL) InsertAtBeg(value int) {
	if l.GetHead() == nil {
		l.SetHead(NewNode(value))
		return
	}
	// point the new node to node pointed by head
	// then point head to the new node
	newNode := NewNode(value)
	newNode.Next = l.GetHead()
	l.SetHead(newNode)
}

func (l *LL) InsertAtPos(value int, pos int) {
	if pos == 0 {
		l.InsertAtBeg(value)
		return
	}
	first := l.GetHead()
	second := l.GetHead().Next
	newNode := NewNode(value)
	s := 1
	for second != nil {
		if s == pos {
			newNode.Next = second
			first.Next = newNode
		}
		first = first.Next
		second = second.Next
		s++
	}
}
