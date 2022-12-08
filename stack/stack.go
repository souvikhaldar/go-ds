package stack

type Stacker interface {
	Push(val interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
}
