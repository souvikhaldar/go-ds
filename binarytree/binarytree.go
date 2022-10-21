package binarytree

type BinaryTree interface {
	AddNode(*Node) error
	DelNode(value int) error
	Search(value int) bool
}
type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{
		value: v,
		left:  nil,
		right: nil,
	}
}

func (n *Node) SetLeft(node *Node) {
	n.left = node
}

func (n *Node) SetRight(node *Node) {
	n.right = node
}

func (n *Node) SetValue(v int) {
	n.value = v
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetRight() *Node {
	return n.right
}
