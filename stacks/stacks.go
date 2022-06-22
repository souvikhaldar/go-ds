package stacks

import "github.com/souvikhaldar/go-ds/linkedlist"

type Stacks interface {
	Push(val int)
	Pop() (int, error)
	Top() (int, error)
	IsEmpty() bool
	Print()
}

type StacksList struct {
	LinkedList linkedlist.LinkedList
}

func NewStacks() *StacksList {
	return &StacksList{
		LinkedList: linkedlist.NewLinkedList(),
	}
}
