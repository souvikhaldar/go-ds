package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	val, _ := q.Dequeue()
	if val != 1 {
		t.Fatal("Should be 1")
	}
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	val, _ = q.Dequeue()
	if val != 5 {
		t.Fatal("Should be 5")
	}
}
