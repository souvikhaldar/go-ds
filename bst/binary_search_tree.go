package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	node := new(Node)
	node.Value = val
	return node
}

type BST interface {
	Insert(int) error
	Search(int) bool
	Delete(int) error
}
