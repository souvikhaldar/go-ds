package stack

import "github.com/souvikhaldar/go-ds/linkedlist"

type Stack struct {
	Tos *linkedlist.LL
}

func NewStack() *Stack {
	return &Stack{
		Tos: linkedlist.NewLinkedList(),
	}
}

func (s *Stack) Push(val interface{}) {
	s.Tos.InsertAtBeg(val)
}
