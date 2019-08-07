package main

import (
	"avl_tree/tree"
	"os"
)

func main() {
	root := tree.AVLBuild(os.Args[1:])
	tree.AVLDraw(root, os.Stdout)
	tree.Inorder(root, os.Stderr)
}
