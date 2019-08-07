package main

import (
	"avl_tree/tree"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	max := 10

	if len(os.Args) > 1 {
		n, e := strconv.Atoi(os.Args[1])
		if e == nil {
			max = n
		}
	}

	var values []int

	for i := 0; i < max; i++ {
		values = append(values, i)
	}

	rand.Shuffle(len(values), func(i, j int) { values[j], values[i] = values[i], values[j] })

	fmt.Printf("%v\n", values)

	var root *tree.Node
	for _, n := range values {
		root, _ = tree.AVLInsert(root, n)
		if !tree.BSTProperty(root) {
			fmt.Fprintf(os.Stderr, "Inserted %d, no longer has BST property\n", n)
		}
	}

	tree.Inorder(root, os.Stdout)
}
