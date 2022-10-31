package stack

import (
	"testing"

	"github.com/souvikhaldar/go-ds/linkedlist"
)

func TestStack(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	if err == nil {
		t.Fatal("Empty stack error: ", err)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	val, _ := s.Pop()
	if val != 4 {
		t.Fatal("Should be 4 but is: ", val)
	}
}

func TestReverse(t *testing.T) {
	g := linkedlist.NewLinkedList()
	g.InsertAtBeg("k")
	g.InsertAtBeg("i")
	g.InsertAtBeg("v")
	g.InsertAtBeg("u")
	g.InsertAtBeg("o")
	g.InsertAtBeg("s")
	f := NewStack()
	f.Push("k")
	f.Push("i")
	f.Push("v")
	f.Push("u")
	f.Push("o")
	f.Push("s")

	val, err := f.Pop()
	llval, err := g.GetFromPos(0)
	err = g.DeleteByPos(0)
	for err == nil {
		if val != llval {
			t.Fatal("Did not match")
		}
		val, err = f.Pop()
		llval, err = g.GetFromPos(0)
		err = g.DeleteByPos(0)
	}

}
