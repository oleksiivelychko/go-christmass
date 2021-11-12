package tree_traversal

type Node struct {
	value int
	left  *Node
	right *Node
}

func TreeTraversal(t *Node) int {
	var value = t.value

	if t.left != nil {
		value += TreeTraversal(t.left)
	}

	if t.right != nil {
		value += TreeTraversal(t.right)
	}

	return value
}
