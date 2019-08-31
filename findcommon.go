package main

import (
	"avl_tree/tree"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var root *tree.Node
	var a, b int
	for i, str := range os.Args[1:] {
		if str == "--" {
			a, _ = strconv.Atoi(os.Args[i+2])
			b, _ = strconv.Atoi(os.Args[i+3])
			fmt.Fprintf(os.Stderr, "Finding common ancestor of %d and %d (%s and %s)\n", a, b, os.Args[i+1], os.Args[i+2])
			break
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		root = tree.Insert(root, n)
	}

	fmt.Fprintf(os.Stderr, "\n")
	tree.Inorder(root, os.Stderr)
	fmt.Fprintf(os.Stderr, "\n")

	if !tree.BSTProperty(root) {
		fmt.Fprintf(os.Stderr, "Not a binary search tree\n")
		return
	}

	if a > b {
		a, b = b, a
	}

	x, foundit := findcommon(root, a, b)
	if !foundit {
		fmt.Fprintf(os.Stderr, "Did not find common ancestor of %d and %d\n", a, b)
		return
	}
	fmt.Fprintf(os.Stderr, "%d is common ancestor of %d and %d\n", x, a, b)
}

func findcommon(node *tree.Node, a, b int) (int, bool) {
	if node == nil {
		return 0, false
	}
	if a < node.Data && b > node.Data {
		return node.Data, true
	}
	if a < node.Data && b < node.Data {
		return findcommon(node.Left, a, b)
	}
	return findcommon(node.Right, a, b)
}
