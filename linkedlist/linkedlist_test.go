package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	l := NewLinkedList()
	l.InsertAtEnd(1)
	l.InsertAtEnd(2)
	l.InsertAtEnd(3)
	l.InsertAtEnd(5)
	l.InsertAtBeg(0)
	l.InsertAtPos(4, 4)
	l.PrintList()

	data, err := l.GetFromPos(1)
	if err != nil || data != 1 {
		t.Fatal("1 should be found at position 1 but got: ", data)
	}

}
