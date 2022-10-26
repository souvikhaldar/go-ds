package bst

type IterativeBst struct {
	Root *Node
}

func NewIterativeBst() *IterativeBst {
	return &IterativeBst{
		Root: nil,
	}
}

func (ibst *IterativeBst) Insert(node *Node) error {
	if ibst.Root == nil {
		ibst.Root = node
		return nil
	}
	temp := ibst.Root

	for {
		if node.Value <= temp.Value {
			if temp.Left == nil {
				temp.Left = node
				break
			} else {
				temp = temp.Left
			}
		} else {
			if temp.Right == nil {
				temp.Right = node
				break
			} else {
				temp = temp.Right
			}
		}
	}
	return nil
}

func (ibst *IterativeBst) Search(val int) bool {
	if ibst.Root == nil {
		return false
	}
	temp := ibst.Root
	for temp != nil {
		if val == temp.Value {
			return true
		} else if val <= temp.Value {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}
	return false
}

func (ibst *IterativeBst) Delete(int) error {
	return nil
}
