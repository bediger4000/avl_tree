package tree

/*
 * Plain binary search tree insertion
 */

func Insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Data: value, Depth: 0, Left: nil, Right: nil}
	}
	if value < node.Data {
		node.Left = Insert(node.Left, value)
		return node
	}
	node.Right = Insert(node.Right, value)
	return node
}
