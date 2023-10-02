package tree_traversal

import "testing"

func makeTree() *node {
	var tree = &node{value: 0, left: nil, right: nil}
	tree.left = &node{value: 1, left: nil, right: nil}
	tree.left.left = &node{value: 2, left: nil, right: nil}
	tree.left.left.left = &node{value: 3, left: nil, right: nil}
	tree.left.left.right = &node{value: 4, left: nil, right: nil}
	tree.left.right = &node{value: 5, left: nil, right: nil}
	tree.right = &node{value: 6, left: nil, right: nil}

	return tree
}

func TestTreeTraversal(t *testing.T) {
	sum := treeTraversal(makeTree())
	if sum != 21 {
		t.Errorf("sum %d is wrong", sum)
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	tree := makeTree()
	for i := 0; i < b.N; i++ {
		treeTraversal(tree)
	}
}
