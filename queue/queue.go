package queue

type Queuer interface {
	Enqueue(int)
	Dequeue() (int, error)
}
