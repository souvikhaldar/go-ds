package queue

type Queuer interface {
	Enqueue(interface{})
	Dequeue() (interface{}, error)
}
