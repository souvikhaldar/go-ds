package bst

import "fmt"

// Find the node with maximum value
// in a tree with pointer to root stored
// in the arg `root`
// It returns pointer to the node with Maximum value
func FindMaxNodeInTree(root *Node) *Node {
	if root == nil {
		fmt.Println(ErrEmptyTree)
		return nil
	}

	if root.Right == nil {
		return root
	}
	return FindMaxNodeInTree(root.Right)
}

func IsALeafNode(node *Node) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func IsNodeWithOneChild(node *Node) bool {
	if (node.Left == nil && node.Right != nil) || (node.Left != nil && node.Right == nil) {
		return true
	}
	return false
}
