package linkedlist

import "fmt"

func (l *LL) DeleteByPos(pos int) error {
	if l.GetHead() == nil {
		return ErrEmptyList
	}
	// delete the 0th node
	if pos == 0 {
		l.SetHead(l.GetHead().GetNext())
		return nil
	}

	// deleting non-zeroth node
	first := l.GetHead()
	second := l.GetHead().GetNext()
	for i := 1; second != nil; i++ {
		if i == pos {
			first.SetNext(second.GetNext())
			return nil
		}
		first = first.GetNext()
		second = second.GetNext()
	}

	return fmt.Errorf("Node not found")
}

func (l *LL) DeleteByValue(value interface{}) error {
	if l.GetHead() == nil {
		return ErrEmptyList
	}
	// delete the 0th node
	if l.GetHead().GetValue() == value {
		l.SetHead(l.GetHead().GetNext())
		return nil
	}

	// deleting non-zeroth node
	first := l.GetHead()
	second := l.GetHead().GetNext()
	for i := 1; second != nil; i++ {
		if second.GetValue() == value {
			first.SetNext(second.GetNext())
			return nil
		}
		first = first.GetNext()
		second = second.GetNext()
	}

	return fmt.Errorf("Node not found")

}
