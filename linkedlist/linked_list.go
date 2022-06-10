package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func NewNode(val int) *Node {
	return &Node{
		Value: val,
		Next:  nil,
	}
}

type LinkedList interface {
	// inserting to the list
	InsertAtEnd(value int)
	InsertAtBeg(value int)
	InsertAtPos(value int, pos int)

	// fetch from the list
	GetFromPos(pos int) (int, error)
}

type LL struct {
	head *Node
}

func (l *LL) GetHead() *Node {
	return l.head
}

func (l *LL) SetHead(n *Node) {
	l.head = n
}

func NewLinkedList() *LL {
	return &LL{
		head: nil,
	}
}
