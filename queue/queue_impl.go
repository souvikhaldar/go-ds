package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(val interface{}) *Node {
	return &Node{
		Value: val,
		Next:  nil,
	}
}

type Queue struct {
	Front *Node
	Rear  *Node
}

func NewQueue() *Queue {
	return &Queue{
		Front: nil,
		Rear:  nil,
	}
}

func (q *Queue) Enqueue(val interface{}) {
	newNode := NewNode(val)
	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
		return
	}
	q.Rear.Next = newNode
	q.Rear = newNode
	return
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}

	temp := q.Front
	if q.Front != nil {
		q.Front = q.Front.Next
	}
	return temp.Value, nil
}

func (q *Queue) PrintQueue() {
	temp := q.Front
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}

func (q *Queue) IsEmpty() bool {
	if q.Front == nil {
		return true
	}
	return false
}
