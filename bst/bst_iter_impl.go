package bst

import (
	"errors"
	"fmt"

	"github.com/souvikhaldar/go-ds/queue"
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

// Deleting a node in BST is complicated, there are three cases for it, when the node to be deleted is:
// 	   1. Leaf node: Simply delete the node
// 	   2. Node with one child: Point the parent to the child then delete the node
// 	   3. Node with two children: Follow the following steps
// 		      1. Find the node with maximum value in left subtree or Minimum value in right subtree.
// 		      2. Copy the node's value to the given node
// 		      3. Now the node from which we copied can be either Leaf node or Node with one child because minimum in right tree means it can't have left child as it is then minimum (similar case for maximum in right subtree), hence we can follow either case 1 or case 2 above.

func (ibst *IterativeBst) Delete(val int) error {
	if ibst.Root == nil {
		return ErrEmptyTree
	}
	temp := ibst.Root

	for temp != nil {
		if val == temp.Value {
			// check for the above three cases
			// Case 1. is it a leaf node?
			if temp.Left == nil && temp.Right == nil {
				// then simply delete the node
				// NOTE: making the value at mem pointed
				// by temp pointer as nil will deallocate it
				*temp = Node{}
				return nil
			} else if (temp.Left == nil && temp.Right != nil) || (temp.Right == nil && temp.Left != nil) {
				// Case 2. a node with only one child
				// then copy the value of the child then delete the child
				if temp.Left != nil {
					temp.Value = temp.Left.Value
					temp.Left = nil
					return nil
				} else if temp.Right != nil {
					temp.Value = temp.Right.Value
					temp.Right = nil
					return nil
				}
			} else {
				// Case 3. a node with both children
				// then find node with max value in left subtree
				maxNode := FindMaxNodeInTree(temp.Left)
				temp.Value = maxNode.Value
				*maxNode = Node{}
				return nil
			}
		} else if val > temp.Value {
			temp = temp.Right
		} else {
			temp = temp.Left
		}
	}
	return ErrValNotFound
}

// Find the node with maximum value
// in a tree with pointer to root stored
// in the arg `root`
// It returns pointer to the node with Maximum value
func FindMaxNodeInTree(root *Node) *Node {
	if root.Right == nil {
		return root
	}
	return FindMaxNodeInTree(root.Right)
}

func (ibst *IterativeBst) PrintLevelOrder() {
	q := queue.NewQueue()
	if ibst.Root == nil {
		return
	}
	temp := ibst.Root
	for temp != nil {
		fmt.Println(temp.Value)
		q.Enqueue(temp.Left)
		q.Enqueue(temp.Right)
		for !q.IsEmpty() {
			val, _ := q.Dequeue()
			fmt.Println(val)
		}
	}
}
