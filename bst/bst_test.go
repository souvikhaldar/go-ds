package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := NewIterativeBst()
	bst.Insert(NewNode(2))
	bst.Insert(NewNode(5))
	bst.Insert(NewNode(1))
	bst.Insert(NewNode(8))
	bst.Insert(NewNode(9))
	bst.Insert(NewNode(4))
	bst.Insert(NewNode(6))
	largest := FindMaxNodeInTree(bst.Root).Value
	if largest != 9 {
		t.Fatal("9 is the largest")
	}
	if bst.Search(5) != true {
		t.Fatal("5 is rather present")
	}
	if bst.Search(3) == true {
		t.Fatal("3 is not present")
	}
	//bst.Delete(5)
	//if bst.Search(5) == true {
	//	t.Fatal("5 is not present")
	//}
	bst.PrintLevelOrder()
}
