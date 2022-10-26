package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := NewIterativeBst()
	bst.Insert(NewNode(2))
	bst.Insert(NewNode(5))
	bst.Insert(NewNode(1))
	if bst.Search(5) != true {
		t.Fatal("5 is rather present")
	}
	if bst.Search(3) == true {
		t.Fatal("3 is not present")
	}
}
