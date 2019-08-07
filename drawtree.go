package main

import (
	"avl_tree/tree"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var root *tree.Node
	for _, str := range os.Args[1:] {
		n, e := strconv.Atoi(str)
		if e != nil {
			continue
		}
		root = tree.Insert(root, n)
		if !tree.BSTProperty(root) {
			fmt.Fprintf(os.Stderr, "Inserted %d, no longer has BST property\n", n)
		}
	}

	tree.Draw(root, os.Stdout)
}
