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
		node.Left.Parent = node
		d := node.Left.Depth + 1
		if d > node.Depth {
			node.Depth = d
		}
		return node
	}
	node.Right = Insert(node.Right, value)
	node.Right.Parent = node
	d := node.Right.Depth + 1
	if d > node.Depth {
		node.Depth = d
	}
	return node
}
