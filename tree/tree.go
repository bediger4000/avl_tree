package tree

import (
	"fmt"
	"io"
	"math"
)

type Node struct {
	Data   int
	Depth  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func Inorder(node *Node, out io.Writer) {
	if node == nil {
		return
	}
	Inorder(node.Left, out)
	fmt.Fprintf(out, "%d\n", node.Data)
	Inorder(node.Right, out)
}

// Return true if tree has "Binary seach tree"
// property, false otherwise.
func BSTProperty(root *Node) bool {
	if !bst(root.Left, int(math.MinInt32), root.Data) {
		return false
	}
	if !bst(root.Right, root.Data, int(math.MaxInt32)) {
		return false
	}
	return true
}

// function that actually checks BST property for
// a given node somewhere in the tree.
func bst(node *Node, min int, max int) bool {

	if node == nil {
		return true
	}

	if !(node.Data > min && node.Data < max) {
		return false
	}
	if !bst(node.Left, min, node.Data) {
		return false
	}
	if !bst(node.Right, node.Data, max) {
		return false
	}

	return true
}
