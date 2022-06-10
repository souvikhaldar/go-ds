package linkedlist

import "fmt"

func (l *LL) GetFromPos(pos int) (int, error) {
	temp := l.GetHead()
	for i := 0; temp != nil; i++ {
		if i == pos {
			return temp.Value, nil
		}
		temp = temp.Next
	}
	return 0, fmt.Errorf("position not found")
}
