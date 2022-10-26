package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}

type BST interface {
	Insert(*Node) error
	Search(int) bool
	Delete(int) error
}
