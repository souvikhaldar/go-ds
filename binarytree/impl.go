package binarytree

type BT struct {
	root *Node
}

func NewBinaryTree() BinaryTree {
	return &BT{
		root: nil,
	}
}

func (b *BT) GetRoot() *Node {
	return b.root
}

func (b *BT) SetRoot(r *Node) {
	b.root = r
}

func (b *BT) AddNode(n *Node) error {
	temp := b.GetRoot()
	if temp == nil {
		b.SetRoot(n)
	}
	for temp != nil {
		if temp.GetLeft() != nil {
			temp = temp.GetLeft()
			continue
		}
		if temp.GetRight() != nil {
			temp = temp.GetRight()
			continue
		}
	}
	temp.SetLeft(n)
	return nil
}

func (b *BT) DelNode(value int) error {
	panic("not implemented") // TODO: Implement
}

func (b *BT) Search(value int) bool {
	return search(b.GetRoot(), value)
}

func search(n *Node, val int) bool {
	found := false
	if n == nil {
		return found
	}
	if n.GetValue() == val {
		found = true
		return found
	}
	search(n.GetLeft(), val)
	search(n.GetRight(), val)
	return found
}
