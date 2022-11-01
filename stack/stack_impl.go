package stack

import (
	"errors"

	"github.com/souvikhaldar/go-ds/linkedlist"
)

var (
	ErrEmptyStack = errors.New("Stack is empty")
)

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

func (s *Stack) Pop() (interface{}, error) {
	if s.Tos == nil {
		return nil, ErrEmptyStack
	}
	val, _ := s.Tos.GetFromPos(0)
	err := s.Tos.DeleteByPos(0)
	return val, err
}

func (s *Stack) IsEmpty() bool {
	return s.Tos == nil
}
