package tree_traversal

import (
	"testing"
)

func initTree() *Node {
	var tree = &Node{0, nil, nil}
	tree.left = &Node{1, nil, nil}
	tree.left.left = &Node{2, nil, nil}
	tree.left.left.left = &Node{3, nil, nil}
	tree.left.left.right = &Node{4, nil, nil}
	tree.left.right = &Node{5, nil, nil}
	tree.right = &Node{6, nil, nil}
	return tree
}

func TestNode(t *testing.T) {
	node := &Node{0, &Node{1, nil, nil}, &Node{2, nil, nil}}
	if node.left.value != 1 {
		t.Errorf("[type Node struct] -> %d != 1", node.left.value)
	}
	if node.right.value != 2 {
		t.Errorf("[type Node struct] -> %d != 2", node.right.value)
	}
}

func TestTreeTraversal(t *testing.T) {
	result := TreeTraversal(initTree())
	if result != 21 {
		t.Errorf("[func TreeTraversal(t *Node) int] -> %d != 21", result)
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	tree := initTree()
	for i := 0; i < b.N; i++ {
		TreeTraversal(tree)
	}
}
