package main

import (
	"avl_tree/tree"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	RandomAVLInsert = 0
	RandomBSTInsert = iota
)

func main() {
	rand.Seed(time.Now().UnixNano())

	insertMethod := flag.String("I", "AVL", "Insert method, AVL or BST")
	insertCount := flag.Int("n", 10, "Count of integers to insert")
	insertFile := flag.String("f", "", "File name to write GraphViz output")
	flag.Parse()

	var values []int

	for i := 0; i < *insertCount; i++ {
		values = append(values, i)
	}

	insert := tree.AVLInsert
	if *insertMethod == "BST" {
		insert = insertFacade
	}

	rand.Shuffle(len(values), func(i, j int) { values[j], values[i] = values[i], values[j] })

	var root *tree.Node
	for _, n := range values {
		root, _ = insert(root, n)
		if !tree.BSTProperty(root) {
			fmt.Fprintf(os.Stderr, "Inserted %d, no longer has BST property\n", n)
		}
	}

	fout := os.Stdout
	if *insertFile != "" {
		var err error
		fout, err = os.Create(*insertFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	tree.Draw(root, fout)

	fout.Close()
}

func insertFacade(node *tree.Node, value int) (*tree.Node, int) {
	return tree.Insert(node, value), 0
}
