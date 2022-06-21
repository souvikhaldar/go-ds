package stacks

import "github.com/souvikhaldar/go-ds/linkedlist"

type Stacks interface {
	Push(val int)
	Pop() (int, error)
	Top() (int, error)
	IsEmpty() bool
}

type StacksList struct {
	LinkedList linkedlist.LinkedList
}
