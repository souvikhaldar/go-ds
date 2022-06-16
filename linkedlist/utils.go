package linkedlist

func (l *LL) Reverse() {
	first := l.GetHead()
	var second *Node
	var third *Node

	if first != nil {
		second = first.GetNext()
	} else {
		// since there is no element
		return
	}

	if second != nil {
		third = second.GetNext()
	} else {
		// since there is only one element
		// so it is same as reverse of it
		return
	}
	if third == nil {
		// only two nodes in the list
		second.SetNext(first)
		l.SetHead(second)
		first.SetNext(nil)
	}

	// for >= 3 nodes in the list

	// make the first node point to nil
	// i.e the last node
	first.SetNext(nil)
	for third != nil {
		second.SetNext(first)

		// move the pointers by one position
		first = second
		second = third
		third = third.GetNext()
	}
	second.SetNext(first)
	l.SetHead(second)
}

func (l *LL) ReverseRecursively(node *Node) {
	if node.GetNext() == nil {
		// the last node

		// make the head point to this node
		l.SetHead(node)
		return
	}
	l.ReverseRecursively(node.GetNext())
	node.GetNext().SetNext(node)
	node.SetNext(nil)
	return
}
