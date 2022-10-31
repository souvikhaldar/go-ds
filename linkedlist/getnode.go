package linkedlist

import "fmt"

func (l *LL) GetFromPos(pos int) (interface{}, error) {
	temp := l.GetHead()
	for i := 0; temp != nil; i++ {
		if i == pos {
			return temp.GetValue(), nil
		}
		temp = temp.GetNext()
	}
	return nil, fmt.Errorf("position not found")
}
