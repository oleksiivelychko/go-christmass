package treetraversal

type node struct {
	value int
	left  *node
	right *node
}

func traverse(t *node) int {
	var sum = t.value

	if t.left != nil {
		sum += traverse(t.left)
	}

	if t.right != nil {
		sum += traverse(t.right)
	}

	return sum
}
