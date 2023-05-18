package tree_traversal

import "testing"

func makeTree() *Node {
	var tree = &Node{Value: 0, Left: nil, Right: nil}
	tree.Left = &Node{Value: 1, Left: nil, Right: nil}
	tree.Left.Left = &Node{Value: 2, Left: nil, Right: nil}
	tree.Left.Left.Left = &Node{Value: 3, Left: nil, Right: nil}
	tree.Left.Left.Right = &Node{Value: 4, Left: nil, Right: nil}
	tree.Left.Right = &Node{Value: 5, Left: nil, Right: nil}
	tree.Right = &Node{Value: 6, Left: nil, Right: nil}

	return tree
}

func TestTreeTraversal(t *testing.T) {
	sum := TreeTraversal(makeTree())
	if sum != 21 {
		t.Errorf("sum %d is wrong", sum)
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	tree := makeTree()
	for i := 0; i < b.N; i++ {
		TreeTraversal(tree)
	}
}
