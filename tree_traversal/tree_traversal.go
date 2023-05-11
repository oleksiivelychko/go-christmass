package tree_traversal

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func TreeTraversal(tree *Node) int {
	var value = tree.Value

	if tree.Left != nil {
		value += TreeTraversal(tree.Left)
	}

	if tree.Right != nil {
		value += TreeTraversal(tree.Right)
	}

	return value
}
