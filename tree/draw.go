package tree

import (
	"fmt"
	"io"
)

func Draw(root *Node, out io.Writer) {
	fmt.Printf("digraph g {\n")
	preorderDraw(root, out)
	fmt.Printf("\n}\n")
}

// Does a pre-order traverse of tree so as to
// print out the nodes before it prints any edge
// referencing a node. Not sure this is necessary,
// but GraphViz is a little obtuse.
func preorderDraw(node *Node, out io.Writer) {

	fmt.Fprintf(out, "N%p [label=\"%d\"];\n", node, node.Data)

	if node.Left != nil {
		// Draw edges with min:max that every node below has to be between
		preorderDraw(node.Left, out)
		fmt.Fprintf(out, "N%p -> N%p;\n", node, node.Left)
	} else {
		// Draw a solid dot node to keep dot honest about left and right children
		fmt.Fprintf(out, "N%pL [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pL;\n", node, node)
	}

	if node.Right != nil {
		preorderDraw(node.Right, out)
		fmt.Fprintf(out, "N%p -> N%p;\n", node, node.Right)
	} else {
		fmt.Fprintf(out, "N%pR [shape=point];\n", node)
		fmt.Fprintf(out, "N%p -> N%pR;\n", node, node)
	}
}
