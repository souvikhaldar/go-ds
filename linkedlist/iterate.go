package linkedlist

import "fmt"

func (l *LL) Iterate() []int {
	out := make([]int, 0)
	temp := l.GetHead()
	for temp != nil {
		out = append(out, temp.GetValue())
		temp = temp.GetNext()
	}
	return out
}

func (l *LL) PrintList() {
	for _, v := range l.Iterate() {
		fmt.Println(v)
	}
}
