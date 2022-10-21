package binarytree

import "testing"

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree()
	bt.AddNode(NewNode(1))
	bt.AddNode(NewNode(2))
	bt.AddNode(NewNode(3))
	bt.AddNode(NewNode(4))
	bt.AddNode(NewNode(5))
	if bt.Search(2) != true {
		t.Fatal("2 is present")
	}
	if bt.Search(6) == true {
		t.Fatal("2 is present")
	}
}
