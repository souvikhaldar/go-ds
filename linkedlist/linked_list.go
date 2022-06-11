package linkedlist

type Node struct {
	value int
	next  *Node
}

func NewNode(val int) *Node {
	return &Node{
		value: val,
		next:  nil,
	}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(nxt *Node) {
	n.next = nxt
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) SetValue(val int) {
	n.value = val
}

type LinkedList interface {
	// inserting to the list
	InsertAtEnd(value int)
	InsertAtBeg(value int)
	InsertAtPos(value int, pos int)

	// fetch from the list
	GetFromPos(pos int) (int, error)

	// delete from the list
	DeleteByPos(pos int) error
	DeleteByValue(value int) error
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
