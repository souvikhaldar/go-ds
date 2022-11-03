package bst

import (
	"errors"
	"fmt"

	"github.com/souvikhaldar/go-ds/queue"
	"github.com/souvikhaldar/go-ds/stack"
)

var (
	ErrEmptyTree   = errors.New("tree is empty")
	ErrValNotFound = errors.New("Value not found")
)

const (
	MinInt = -9223372036854775808
)

type IterativeBst struct {
	Root *Node
}

func NewIterativeBst() *IterativeBst {
	return &IterativeBst{
		Root: nil,
	}
}

func (ibst *IterativeBst) Insert(val int) error {
	node := NewNode(val)
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

// Deleting a node in BST is complicated, there are three cases for it, when the node to be deleted is:
// 	   1. Leaf node: Simply delete the node
// 	   2. Node with one child: Point the parent to the child then delete the node
// 	   3. Node with two children: Follow the following steps
// 		      1. Find the node with maximum value in left subtree or Minimum value in right subtree.
// 		      2. Copy the node's value to the given node
// 		      3. Now the node from which we copied can be either Leaf node or Node with one child because minimum in right tree means it can't have left child as it is then minimum (similar case for maximum in right subtree), hence we can follow either case 1 or case 2 above.

func (ibst *IterativeBst) Delete(val int) error {
	parent := ibst.Root
	temp := ibst.Root
	if temp == nil {
		return ErrEmptyTree
	}
	for temp != nil {
		if val == temp.Value {
			// check for the above three cases
			// Case 1. is it a leaf node?
			if IsALeafNode(temp) {
				// then simply delete the node
				if parent.Left.Value == temp.Value {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				return nil
			} else if IsNodeWithOneChild(temp) {
				// Case 2. a node with only one child

				// bypas the connection to the child
				if temp.Left != nil {
					// relation of parent to temp?
					if parent.Left.Value == temp.Value {
						parent.Left = temp.Left
					} else if parent.Right.Value == temp.Value {
						parent.Right = temp.Left
					}
				} else if temp.Right != nil {
					// relation of parent to temp?
					if parent.Left.Value == temp.Value {
						parent.Left = temp.Right
					} else if parent.Right.Value == temp.Value {
						parent.Right = temp.Right
					}
				}
				return nil
			} else {
				// Case 3. a node with both children
				// then find node with max value in left subtree and replace the value
				maxNode := FindMaxNodeInTree(temp.Left)
				temp.Value = maxNode.Value
				maxNode.Value = val

				// proceed the search to left subtree because
				// we used max in left subtree for replacement
				parent = temp
				temp = temp.Left
			}

		} else if val > temp.Value {
			parent = temp
			temp = temp.Right
		} else {
			parent = temp
			temp = temp.Left
		}
	}
	return ErrValNotFound
}

func (ibst *IterativeBst) PrintLevelOrder() {
	q := queue.NewQueue()
	if ibst.Root == nil {
		return
	}
	q.Enqueue(ibst.Root)
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
			return
		}
		node, ok := val.(*Node)
		if !ok {
			fmt.Println("Unable to assert to type *Node")
			return
		}
		fmt.Println(node.Value)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
}

// PreOrder prints in the order of root, left then right
func (ibst *IterativeBst) PreOrder() {
	if ibst.Root == nil {
		return
	}
	temp := ibst.Root
	s := stack.NewStack()
	for temp != nil {
		fmt.Println(temp.Value)
		if temp.Right != nil {
			s.Push(temp.Right)
		}

		if temp.Left == nil {
			t, _ := s.Pop()
			var ok bool
			temp, ok = t.(*Node)
			if !ok {
				return
			}
		} else {
			temp = temp.Left
		}
	}
}

// InOrder prints left, root, right
func (ibst *IterativeBst) InOrder() {
	if ibst.Root == nil {
		return
	}
	temp := ibst.Root
	s := stack.NewStack()

	for temp != nil || !s.IsEmpty() {
		if temp == nil {
			t, _ := s.Pop()
			if t == nil {
				return
			}
			temp, _ = t.(*Node)
			fmt.Println(temp.Value)
			temp = temp.Right
			continue
		}
		s.Push(temp)
		temp = temp.Left
	}
}
