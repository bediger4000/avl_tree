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
			fmt.Fprintf(os.Stderr, "Finding common ancestor of %d and %d (%s and %s)\n", a, b, os.Args[i+2], os.Args[i+3])
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

	aNode := findValueNode(root, a)
	if aNode == nil {
		fmt.Fprintf(os.Stderr, "No value %d in tree\n", a)
		return
	}

	aLen := 0
	for n := aNode; n != nil; n = n.Parent {
		aLen++
		fmt.Fprintf(os.Stderr, "%d/%d -> ", n.Data, n.Depth)
	}
	fmt.Fprintf(os.Stderr, "\n%d node path length %d\n", aNode.Data, aLen)

	bNode := findValueNode(root, b)
	if bNode == nil {
		fmt.Fprintf(os.Stderr, "No value %d in tree\n", b)
		return
	}
	bLen := 0
	for n := bNode; n != nil; n = n.Parent {
		bLen++
		fmt.Fprintf(os.Stderr, "%d/%d -> ", n.Data, n.Depth)
	}
	fmt.Fprintf(os.Stderr, "\n%d node path length %d\n", bNode.Data, bLen)

}

func findValueNode(node *tree.Node, value int) *tree.Node {
	if node == nil {
		return nil
	}
	if node.Data == value {
		return node
	}
	if value < node.Data {
		return findValueNode(node.Left, value)
	}
	if value > node.Data {
		return findValueNode(node.Right, value)
	}
	return nil
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
