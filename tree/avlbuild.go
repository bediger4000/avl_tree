package tree

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 * Convenience function: build a AVL tree from a []string.
 * Assumes most or all elements of []string constitute
 * representations of integers.
 * Checks BST property after each insertion.
 */

func AVLBuild(args []string) (root *Node) {
	for _, str := range args {
		value, err := strconv.Atoi(str)
		if err != nil {
			log.Print(err)
			continue
		}
		root, _ = AVLInsert(root, value)
		if !BSTProperty(root) {
			fmt.Fprintf(os.Stderr, "Inserted %d, no longer has BST property\n", value)
		}
	}
	return root
}
