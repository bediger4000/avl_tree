package tree

/*
 * GraphViz tree drawing, with per-node annotation
 * about conforming to Binary Search Tree property
 */

import (
	"fmt"
	"io"
)

func BSTDraw(root *Node, out io.Writer) {
	fmt.Printf("digraph g {\n")
	bstPreorderDraw(root, -1, 9999, out)
	fmt.Printf("\n}\n")
}

// Does a pre-order traverse of tree so as to
// print out the nodes before it prints any edge
// referencing a node. Not sure this is necessary,
// but GraphViz is a little obtuse.
func bstPreorderDraw(node *Node, min int, max int, out io.Writer) {

	// Label each node with conformance to BST property
	conforms := node.Data > min && node.Data < max
	fmt.Fprintf(out, "N%p [label=\"%d/%v\"];\n", node, node.Data, conforms)

	if node.Left != nil {
		// Draw edges with min:max that every node below has to be between
		bstPreorderDraw(node.Left, min, node.Data, out)
		fmt.Fprintf(out, "N%p -> N%p [label=\"%d:%d\"];\n", node, node.Left, min, node.Data)
	} else {
		// Draw a solid dot node to keep dot honest about left and right children
		fmt.Fprintf(out, "N%pL [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pL;\n", node, node)
	}

	if node.Right != nil {
		bstPreorderDraw(node.Right, node.Data, max, out)
		fmt.Fprintf(out, "N%p -> N%p [label=\"%d:%d\"];\n", node, node.Right, node.Data, max)
	} else {
		fmt.Fprintf(out, "N%pR [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pR;\n", node, node)
	}
}
