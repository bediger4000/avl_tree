package tree

/*
 * GraphViz rendition of a what should constitute a binary
 * search tree, with depth (leaf == 0) and "balance factor"
 * (left child depth - right child depth) annotations
 */

import (
	"fmt"
	"io"
)

func AVLDraw(root *Node, out io.Writer) {
	fmt.Printf("digraph g {\n")
	avlPreorderDraw(root, out)
	fmt.Printf("\n}\n")
}

// Does a pre-order traverse of tree so as to
// print out the nodes before it prints any edge
// referencing a node. Not sure this is necessary,
// but GraphViz is a little obtuse.
func avlPreorderDraw(node *Node, out io.Writer) {

	hLeft := -1
	if node.Left != nil {
		hLeft = node.Left.Depth
	}
	hRight := -1
	if node.Right != nil {
		hRight = node.Right.Depth
	}
	balanceFactor := hLeft - hRight
	fmt.Fprintf(out, "N%p [label=\"%d,%d/%d\"];\n", node, node.Data, node.Depth, balanceFactor)

	if node.Left != nil {
		// Draw edges with min:max that every node below has to be between
		avlPreorderDraw(node.Left, out)
		fmt.Fprintf(out, "N%p -> N%p;\n", node, node.Left)
	} else {
		// Draw a solid dot node to keep dot honest about left and right children
		fmt.Fprintf(out, "N%pL [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pL;\n", node, node)
	}

	if node.Right != nil {
		avlPreorderDraw(node.Right, out)
		fmt.Fprintf(out, "N%p -> N%p;\n", node, node.Right)
	} else {
		fmt.Fprintf(out, "N%pR [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pR;\n", node, node)
	}
}
