package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := NewIterativeBst()
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(3)
	bst.Insert(6)
	largest := FindMaxNodeInTree(bst.Root).Value
	if largest != 9 {
		t.Fatal("9 is the largest")
	}
	if bst.Search(5) != true {
		t.Fatal("5 is rather present")
	}
	bst.Delete(5)
	if bst.Search(5) == true {
		t.Fatal("5 is not present")
	}
	bst.Delete(2)
	bst.Delete(9)

}

func TestPrint(t *testing.T) {

	bst := NewIterativeBst()
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(-2)
	bst.Insert(-1)
	bst.Insert(14)
	bst.Insert(13)
	//bst.PreOrder()
	//bst.InOrder()
}
