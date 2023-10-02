package tree_traversal

type node struct {
	value int
	left  *node
	right *node
}

func treeTraversal(tree *node) int {
	var value = tree.value

	if tree.left != nil {
		value += treeTraversal(tree.left)
	}

	if tree.right != nil {
		value += treeTraversal(tree.right)
	}

	return value
}
