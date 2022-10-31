package linkedlist

import "errors"

var (
	ErrEmptyList = errors.New("Linked list is empty")
)

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(val interface{}) *Node {
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

func (n *Node) GetValue() interface{} {
	return n.value
}

func (n *Node) SetValue(val interface{}) {
	n.value = val
}

type LinkedList interface {
	GetHead() *Node
	// inserting to the list
	InsertAtEnd(value interface{})
	InsertAtBeg(value interface{})
	InsertAtPos(value interface{}, pos int)

	// fetch from the list
	GetFromPos(pos int) (interface{}, error)

	// delete from the list
	DeleteByPos(pos int) error
	DeleteByValue(value interface{}) error
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
