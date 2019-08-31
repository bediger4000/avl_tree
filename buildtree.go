package main

/*
 * Create an ordinary binary search tree,
 * but print it out with AVL tree annotations
 */

import (
	"avl_tree/tree"
	"os"
)

func main() {
	root := tree.Build(os.Args[1:])
	tree.AVLDraw(root, os.Stdout)
}
