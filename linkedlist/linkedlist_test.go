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

	l.DeleteByPos(0)
	t.Log("After deleting 0th node")
	l.PrintList()
	l.DeleteByValue(3)
	t.Log("After deleting 4th node")
	l.PrintList()

	t.Log("Reversing")
	l.Reverse()
	l.PrintList()

}

func TestReverse(t *testing.T) {
	l := NewLinkedList()
	l.InsertAtEnd(1)
	l.InsertAtEnd(2)
	l.InsertAtEnd(3)
	l.InsertAtEnd(4)
	l.InsertAtEnd(5)

	l.Reverse()
	if l.GetHead().GetValue() != 5 {
		t.Fatal("Iterative reverse failed: got head value ", l.GetHead().GetValue())
	}
	l.ReverseRecursively(l.GetHead())
	if l.GetHead().GetValue() != 1 {
		t.Fatal("Iterative reverse failed: got head value ", l.GetHead().GetValue())
	}

}
