package tree

const (
	InsertNone  = 0
	InsertLeft  = iota
	InsertRight = iota
	InsertLeaf  = iota
)

func AVLInsert(node *Node, value int) (*Node, int) {
	if node == nil {
		return &Node{Data: value, Depth: 0, Left: nil, Right: nil}, InsertLeaf
	}
	var childInsert int
	if value < node.Data {
		node.Left, childInsert = AVLInsert(node.Left, value)
		if childInsert == InsertNone {
			return node, InsertNone
		}
		if childInsert == InsertLeaf {
			node.setDepth()
			return node, InsertLeft
		}

		balanceFactor := node.setDepth()

		if balanceFactor == 2 || balanceFactor == -2 {
			switch childInsert {
			case InsertLeft:
				node = rotateRight(node)
			case InsertRight:
				node.Left = rotateLeft(node.Left)
				node = rotateRight(node)
			}
			return node, InsertNone
		}

		return node, InsertLeft
	}
	node.Right, childInsert = AVLInsert(node.Right, value)
	if childInsert == InsertNone {
		return node, InsertNone
	}
	if childInsert == InsertLeaf {
		node.setDepth()
		return node, InsertRight
	}
	balanceFactor := node.setDepth()
	if balanceFactor == 2 || balanceFactor == -2 {
		switch childInsert {
		case InsertLeft:
			node.Right = rotateRight(node.Right)
			node = rotateLeft(node)
		case InsertRight:
			node = rotateLeft(node)
		}
		return node, InsertNone
	}

	return node, InsertRight
}

/* setDepth() calculates the receiver node's
 * tree depth (leaf == 0) as the max of depths
 * of child node(s).
 * Conveniently calculates and returns the "balance factor"
 * which can trigger a rebalance if it gets > 1 or < -1
 * Calculation complicated by a missing child having a depth
 * of -1.
 */
func (node *Node) setDepth() int {
	rDepth := -1
	node.Depth = 0
	if node.Right != nil {
		node.Depth = node.Right.Depth + 1
		rDepth = node.Right.Depth
	}
	lDepth := -1
	if node.Left != nil {
		lDepth = node.Left.Depth
		if node.Left.Depth >= node.Depth {
			node.Depth = node.Left.Depth + 1
		}
	}
	return lDepth - rDepth
}

func legibleInsert(x int) string {
	switch x {
	case InsertLeft:
		return "left"
	case InsertRight:
		return "right"
	case InsertLeaf:
		return "leaf"
	default:
		return "don't know"
	}
}

func rotateRight(node *Node) *Node {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	node = tmp
	node.Right.setDepth()
	node.setDepth()
	return node
}

func rotateLeft(node *Node) *Node {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	node = tmp
	node.Left.setDepth()
	node.setDepth()
	return node
}
