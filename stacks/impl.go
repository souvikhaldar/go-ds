package stacks

// Push puts an element on top of the stack
// which while implementing using a linked list
// essentially means adding a new node at the beginning
func (s StacksList) Push(val int) {
	s.LinkedList.InsertAtBeg(val)
}

// Pop returns the element present on top of the stack
// and removes it
func (s StacksList) Pop() (int, error) {
	el, err := s.LinkedList.GetFromPos(0)
	if err != nil {
		return el, err
	}
	err := s.LinkedList.DeleteByPos(0)
	return el, err
}

// Top retuns the element from the top of the stack
func (s StacksList) Top() (int, err) {
	return s.LinkedList.GetFromPos(0)
}

func (s StacksList) IsEmpty() bool {
	el, err := s.LinkedList.GetFromPos(0)
	return err != nil
}
