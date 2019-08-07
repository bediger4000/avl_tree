package tree

import (
	"log"
	"strconv"
)

func Build(args []string) (root *Node) {
	for _, str := range args {
		value, err := strconv.Atoi(str)
		if err != nil {
			log.Print(err)
			continue
		}
		root = Insert(root, value)
	}
	return root
}
